package process

import (
	"context"

	"github.com/pkg/errors"
)

type AtomicProcessor interface {
	// Do 方法需要是幂等的
	Do(ctx context.Context) error

	Rollbacker
}

type Rollbacker interface {
	Rollback(ctx context.Context)
	WhenFail(ctx context.Context, err error)
}

// AtomicRun 当其中一个Processor涉及到网络操作，也就是出现了分布式事务时使用
//
// 如果这个方法本身在跑的时候，进程所在机器挂了，怎么办呢？
//
// 	可以考虑将这个方法放在消息队列的消费逻辑里跑，即使挂了，由于消息队列里的消息没被置为完成，下次还可以继续拿到消息来执行。
//
// Processor里的Do方法需要是幂等的，确保多次执行也能得到相同结果
func AtomicRun(ctx context.Context, p AtomicProcessor, ps ...AtomicProcessor) error {
	// 回滚链，p成功之后添加rollback到其里面
	var rbl = make(RollbackerList, 0, len(ps)+1)

	// 执行1
	if err := WrapRecover(p.Do)(ctx); err != nil {
		return err
	}
	rbl = append(rbl, p)

	// 1已经成功了
	for _, p2 := range ps {
		// 执行2
		if err := WrapRecover(p2.Do)(ctx); err != nil {

			// 2执行失败，需要将1回滚
			rbl.Rollback(ctx)

			// 任意一个失败了，执行回滚后，马上返回
			return err
		}
		rbl = append(rbl, p2)
	}

	return nil
}

func WrapRecover(f Func) Func {
	return Func(func(ctx context.Context) (err error) {
		defer func() {
			if r := recover(); r != nil {
				innerErr, ok := r.(error)
				if !ok {
					innerErr = errors.Errorf("recover: %+v", r)
				}
				err = innerErr
			}
		}()
		if err = f(ctx); err != nil {
			return
		}

		return
	})
}

type Func func(ctx context.Context) error

type RollbackerList []Rollbacker

func (l RollbackerList) Rollback(ctx context.Context) {
	for _, rollbacker := range l {
		if err := WrapRecover(func(ctx context.Context) error {

			rollbacker.Rollback(ctx)

			return nil
		})(ctx); err != nil {
			rollbacker.WhenFail(ctx, err)
		}
	}
}
