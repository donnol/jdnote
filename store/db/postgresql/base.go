package pg

import (
	"github.com/jmoiron/sqlx"
)

// TODO:这个值可以从环境变量获取
var isUnitTest bool

// 用于单元测试的全局事务
var globalTx = func() *sqlx.Tx {
	if isUnitTest {
		tx, err := defaultDB.Beginx()
		if err != nil {
			panic(err)
		}
		return tx
	}

	return nil
}()

// New 新建
func New() DB {
	if isUnitTest {
		return globalTx
	}

	return defaultDB
}

// WithTx 事务
func WithTx(f func(tx DB) error) error {
	var tx *sqlx.Tx
	var err error

	if isUnitTest {
		tx = globalTx
	} else {
		tx, err = defaultDB.Beginx()
		if err != nil {
			return err
		}
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

		// 如果是单元测试，直接返回，在单元测试结束时回滚
		if isUnitTest {
			return nil
		}

		err = tx.Commit() // 成功则提交
		if err != nil {
			return err
		}

		return nil
	}

	return err
}
