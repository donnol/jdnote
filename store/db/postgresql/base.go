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

// Base 基底
type Base struct {
	// 直接内嵌接口，虽然可以在调用的时候少写一个字段名，但同时会有一个不好的地方，就是结构体不小心也实现了这些接口的方法，导致使用者调用的方法不是想要的方法，虽然也可以通过在调用的时候主动将字段名写上去，但是还是难以避免这种情况发生时，使用者感到迷茫……
	// 或许可以折衷一下，好像DB这种，里面的方法名很容易在实践中被重写，我们就用指定字段名的方式来写，而Logger这种，里面的方法名并不容易被重写，而且，就算被重写了也不会导致出错，就还是直接内嵌

	// 日志
	utillog.Logger `json:"-" db:"-"`

	// DB
	DB    DB `json:"-" db:"-"`
	oldDB DB `json:"-" db:"-"` // 开始事务的时候，用这个字段保存旧DB值，等事务完成(回滚/提交)后再将它赋回给DB字段，然后将它置空
}

// SetTx 设置事务
func (b *Base) SetTx(tx DB) *Base {
	b.oldDB = b.DB
	b.DB = tx
	return b
}

// New 新建
func (b *Base) New() DB {
	if isUnitTest {
		return globalTx
	}

	return defaultDB
}

// WithTx 事务
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

		b.DB = b.oldDB
		b.oldDB = nil
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

// InjectTx 注入事务
func (b *Base) InjectTx(v interface{}, f func(v interface{}) error) error {
	if err := b.WithTx(func(tx DB) error {
		var err error

		// 注入tx
		v, err = initParamWithDB(v, tx, true)
		if err != nil {
			return err
		}

		// 执行
		err = f(v)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// InjectTx2 注入事务
func (b *Base) InjectTx2(v interface{}, f func(v interface{}) error) error {
	var err error

	// 执行事务
	if err = b.WithTx(func(tx DB) error {
		var errInner error

		// 注入tx
		v, errInner = initParamWithDB(v, tx, true)
		if errInner != nil {
			return errInner
		}

		// 执行
		errInner = f(v)
		if errInner != nil {
			return errInner
		}

		return nil
	}); err != nil {
		return err
	}

	return err
}
