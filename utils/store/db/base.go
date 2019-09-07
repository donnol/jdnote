package db

import (
	"github.com/jmoiron/sqlx"
)

// Base 基底
type Base struct {
	db *sqlx.DB
}

// New 新建
func New(db *sqlx.DB) *Base {
	return &Base{
		db: db,
	}
}

// DB 获取DB
func (b *Base) DB() DB {
	return b.db
}

// WithTx 事务-这种写法必须确定f函数里是否也调用了WithTx，如果不确定，有可能导致事务重复开启，从而出错。所以，在使用唯一一层事务时才使用这个方法
// 适合在最外层使用
func (b *Base) WithTx(f func(tx DB) error) error {
	var tx *sqlx.Tx
	var err error

	tx, err = b.db.Beginx()
	if err != nil {
		return err
	}

	var success bool // 调用f时如果出现panic，err则会无法正常赋值，因此需要此变量
	defer func() {
		if !success {
			tx.Rollback() // 执行f时出现任何问题，都要Rollback
		}
	}()

	err = f(tx)
	if err == nil {
		success = true

		err = tx.Commit() // 成功则提交
		if err != nil {
			return err
		}
		return nil
	}

	return err
}
