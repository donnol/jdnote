package models

import (
	"database/sql"

	"github.com/donnol/jdnote/utils/context"
	utillog "github.com/donnol/jdnote/utils/log"
	"github.com/donnol/jdnote/utils/store/db"
	"github.com/donnol/jdnote/utils/worker"
)

// Base 基底
type Base struct {
	*db.Base
}

// NewBase 新建
func NewBase() *Base {
	return &Base{
		Base: db.New(defaultDB),
	}
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
	Entity func() Entity
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
	w := worker.New(8)
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
	entity := opt.Entity()
	for rows.Next() {
		err := entity.Scan(ctx, rows)
		if err != nil {
			return err
		}
		acNum++

		// 够一批
		if acNum == opt.N {
			w.Push(worker.MakeJob(worker.Do(func() error {
				// 处理数据
				if err := entity.Do(ctx); err != nil {
					return err
				}

				return nil
			}), 0, worker.ErrorHandler(func(err error) {
				utillog.Errorf("do failed, err is %+v", err)
			})))

			acNum = 0
			entity = opt.Entity()
		}
	}
	// TODO:这里不好，要再调一次Do
	// 或者最后几个
	if err := entity.Do(ctx); err != nil {
		return err
	}

	rerr := rows.Close()
	if rerr != nil {
		return err
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		return err
	}

	// 完成
	w.Stop()

	return nil
}

// DB DB接口
type DB = db.DB
