package pg

import (
	"os"

	"github.com/jmoiron/sqlx"
)

const (
	// 单元测试环境变量
	unitTestEnv = "UNIT_TEST_ENV"
)

// 这个值可以从环境变量获取
var isUnitTest = func() bool {
	env, ok := os.LookupEnv(unitTestEnv)
	if !ok || env == "" {
		return false
	}
	return true
}()

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

// Base 基底
type Base struct {
}

// New 新建
func (b *Base) New() DB {
	if isUnitTest {
		return globalTx
	}

	return defaultDB
}

// WithTx 事务-这种写法必须确定f函数里是否也调用了WithTx，如果不确定，有可能导致事务重复开启，从而出错。所以，在使用唯一一层事务时才使用这个方法
// 适合在最外层使用
func (b *Base) WithTx(f func(tx DB) error) error {
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

		// 如果是单元测试，直接回滚并返回
		if isUnitTest {
			err = tx.Rollback()
			if err != nil {
				return err
			}
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
