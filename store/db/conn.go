package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // github.com/lib/pq postgresql驱动
	"github.com/donnol/jdnote/config"
)

// DefaultDB 默认db
var DefaultDB DB

// DB 数据库连接
type DB struct {
	*sqlx.DB
}

// New 新建
func (db *DB) New() *sqlx.DB {
	return DefaultDB.DB
}

// WithTx 事务
func (db *DB) WithTx(f func(tx *sqlx.Tx) error) error {
	tx, err := DefaultDB.DB.Beginx()
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
		return tx.Commit() // 成功则提交
	}

	return err
}

func init() {
	db, err := sqlx.Open(config.DefaultConfig.DB.Scheme, config.DefaultConfig.DB.String())
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	DefaultDB.DB = db
}
