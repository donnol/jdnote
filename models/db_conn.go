package model

import (
	"github.com/donnol/jdnote/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // github.com/lib/pq postgresql驱动
)

const (
	// 单元测试环境变量
	unitTestEnv = "UNIT_TEST_ENV"
)

// defaultDB 默认db
var defaultDB = func() *sqlx.DB {

	db, err := sqlx.Open(config.Default().DB.Scheme, config.Default().DB.String())
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}()

// // 这个值可以从环境变量获取
// var isUnitTest = func() bool {
// 	env, ok := os.LookupEnv(unitTestEnv)
// 	if !ok || env == "" {
// 		return false
// 	}
// 	return true
// }()

// // 用于单元测试的全局事务
// var globalTx = func() *sqlx.Tx {
// 	if isUnitTest {
// 		tx, err := defaultDB.Beginx()
// 		if err != nil {
// 			panic(err)
// 		}
// 		return tx
// 	}

// 	return nil
// }()
