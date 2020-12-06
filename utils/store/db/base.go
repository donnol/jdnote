package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

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
