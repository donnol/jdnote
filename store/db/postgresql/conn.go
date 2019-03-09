package pg

import (
	"github.com/donnol/jdnote/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // github.com/lib/pq postgresql驱动
)

// defaultDB 默认db
var defaultDB *sqlx.DB

func init() {
	db, err := sqlx.Open(config.DefaultConfig.DB.Scheme, config.DefaultConfig.DB.String())
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	defaultDB = db
}

// New 新建
func New() DB {
	return defaultDB
}

// WithTx 事务
func WithTx(f func(tx DB) error) error {
	tx, err := defaultDB.Beginx()
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
