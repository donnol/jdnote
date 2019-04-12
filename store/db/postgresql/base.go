package pg

import (
	"os"

	utillog "github.com/donnol/jdnote/utils/log"
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

// Base 基础
type Base struct {
	DB              `json:"-" db:"-"`
	*utillog.Logger `json:"-" db:"-"`
}

// SetTx 设置事务
func (b *Base) SetTx(tx DB) *Base {
	b.DB = tx
	return b
}

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
