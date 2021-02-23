package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type (
	TxKeyType string
)

const (
	TxKey TxKeyType = "Tx"
)

// DBFromCtxValue 从ctx value获取tx，如果不存在则返回db
func DBFromCtxValue(ctx context.Context, db DB) DB {
	tx, ok := ctx.Value(TxKey).(DB)
	if ok {
		if tx == nil {
			panic(fmt.Errorf("tx from ctx is nil"))
		}
		return tx
	}
	return db
}

func MustGetDBFromCtxValue(ctx context.Context) DB {
	db := DBFromCtxValue(ctx, nil)
	if db == nil {
		panic(fmt.Errorf("can't get tx from ctx"))
	}
	return db
}

func WithTx(db DB, f func(tx DB) error) error {
	var tx *sqlx.Tx
	var err error

	switch dbv := db.(type) {
	case *sqlx.DB:
		tx, err = dbv.Beginx()
		if err != nil {
			return err
		}
	case *sqlx.Tx:
		tx = dbv
	default:
		return fmt.Errorf("Bad DB param: %+v", db)
	}

	var success bool // 调用f时如果出现panic，err则会无法正常赋值，因此需要此变量
	defer func() {
		if !success { // 执行f时出现任何问题，都要Rollback
			if err := tx.Rollback(); err != nil {
				fmt.Printf("Rollback failed, err: %+v\n", err)
			}
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
