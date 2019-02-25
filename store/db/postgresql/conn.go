package pg

import (
	"log"

	"github.com/donnol/jdnote/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // github.com/lib/pq postgresql驱动
)

// defaultDB 默认db
var defaultDB DB

// 单元测试
var (
	// 是否开启
	IsUnitTest bool
	// 测试用事务-测试完成后调用Rollback将事务回滚，可以在测试里添加TestMain，然后在里面defer关闭
	unitTestTx *sqlx.Tx
)

// DB 数据库连接
type DB struct {
	*sqlx.DB
}

// New 新建
func (db *DB) New() *sqlx.DB {
	// TODO:单元测试时，使用一个全局事务，但是事务类型是sqlx.Tx，又不是sqlx.DB
	if IsUnitTest {
		return defaultDB.DB
	}

	return defaultDB.DB
}

// WithTx 事务
func (db *DB) WithTx(f func(tx *sqlx.Tx) error) error {
	tx, err := defaultDB.DB.Beginx()
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

		// 单元测试时，一律回滚事务
		if IsUnitTest {
			err = tx.Rollback()
		} else {
			err = tx.Commit() // 成功则提交
		}
		if err != nil {
			return err
		}

		return nil
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

	defaultDB.DB = db
}
