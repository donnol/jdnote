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

	N    int                                                   // 批数量
	Scan func(context.Context, *sql.Rows) (interface{}, error) // 扫描
	Do   func(context.Context, []interface{}) error            // 处理
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
	var data = make([]interface{}, 0, opt.N)
	for rows.Next() {
		tmp, err := opt.Scan(ctx, rows)
		if err != nil {
			return err
		}
		data = append(data, tmp)
		acNum++

		if acNum == opt.N {
			tmpData := data
			w.Push(worker.MakeJob(worker.Do(func() error {
				// 处理数据
				if err := opt.Do(ctx, tmpData); err != nil {
					return err
				}

				return nil
			}), 0, worker.ErrorHandler(func(err error) {
				utillog.Errorf("do failed, err is %+v", err)
			})))

			data = make([]interface{}, 0, opt.N)
			acNum = 0
		}
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
