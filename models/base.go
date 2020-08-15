package models

import (
	"database/sql"
	"fmt"
	"runtime"

	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/store/db"
	"github.com/donnol/tools/worker"
	"github.com/pkg/errors"
)

// Base 基底
type Base struct {
}

// NewBase 新建
func NewBase() *Base {
	return &Base{}
}

// ProcessOption 选项
type ProcessOption struct {
	Query string
	Args  []interface{}

	N int // 批数量

	// 用函数的话就必须返回interface{}，然后再传进去Do方法里，这样就有点分裂了
	// Scan func(context.Context, *sql.Rows) (interface{}, error) // 扫描
	// Do   func(context.Context, []interface{}) error            // 处理

	// 如果用接口呢？就可以不用将值返回，而是存在实体里，这样实体调用Do方法的时候就不需要传入Scan获取到的返回值
	// 这样也不会有分裂出现
	NewEntity func() Entity
}

// Entity 实体
type Entity interface {
	Scanner
	Doer
}

// Scanner 扫描
type Scanner interface {
	Scan(context.Context, *sql.Rows) error
}

// Doer 处理
type Doer interface {
	Do(context.Context) error
}

// ProcessConcurrent 并发处理
func (b *Base) ProcessConcurrent(ctx context.Context, opt ProcessOption) error {
	// 启动worker
	numCPU := runtime.NumCPU()
	ctx.Logger().Debugf("== numCPU: %d\n", numCPU)
	w := worker.New(numCPU)
	w.Start()

	// 语句查询
	rows, err := ctx.DB().QueryContext(ctx, opt.Query, opt.Args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	// 遍历结果
	// 每找到n条记录，传入worker执行
	var acNum int
	entity := opt.NewEntity()
	for rows.Next() {
		err := entity.Scan(ctx, rows)
		if err != nil {
			return err
		}
		acNum++

		// 够一批
		if acNum == opt.N {
			// 推入任务队列
			if err := w.Push(makeWorkerJob(ctx, entity)); err != nil {
				return err
			}

			// 重置
			acNum = 0
			entity = opt.NewEntity()
		}
	}
	// 剩下还有，但不足一批：需要额外执行一次
	if acNum != 0 && acNum < opt.N {
		if err := w.Push(makeWorkerJob(ctx, entity)); err != nil {
			return err
		}
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		return err
	}

	// 完成
	w.Stop()

	return nil
}

func makeWorkerJob(ctx context.Context, entity Entity) worker.Job {
	return worker.MakeJob(
		newWorkerDo(ctx, entity),
		0,
		newWorkerErrorHandler(ctx),
	)
}

// 参数传递时会复制一个，在任务里使用
func newWorkerDo(ctx context.Context, entity Entity) worker.Do {
	return func() error {
		// 处理数据
		if err := entity.Do(ctx); err != nil {
			return errors.WithMessage(err, fmt.Sprintf("entity: %+v", entity))
		}

		return nil
	}
}

func newWorkerErrorHandler(ctx context.Context) worker.ErrorHandler {
	return func(err error) {
		ctx.Logger().Errorf("do failed, err is %+v", err)
	}
}

// DB DB接口
type DB = db.DB
