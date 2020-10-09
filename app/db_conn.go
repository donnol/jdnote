package app

// const (
// 	// 单元测试环境变量
// 	unitTestEnv = "UNIT_TEST_ENV"
// )

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

// var _ = globalTx
