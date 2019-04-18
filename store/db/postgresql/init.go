package pg

import (
	"reflect"

	"github.com/donnol/jdnote/utils/reflectx"
)

// InitParamWithDB 初始化param里的DB
func InitParamWithDB(param interface{}) (interface{}, error) {
	param, err := initParamWithDB(param, (&Base{}).New(), false)
	if err != nil {
		return param, err
	}

	return param, nil
}

// initParamWithDB 初始化param里的DB
func initParamWithDB(param interface{}, db DB, copy bool) (interface{}, error) {
	// db类型和值
	dbType := reflect.TypeOf((*DB)(nil)).Elem()
	dbValue := reflect.ValueOf(db)

	// 注入
	var err error
	param, err = reflectx.InitParam(param, dbType, dbValue, copy)
	if err != nil {
		return param, err
	}

	return param, nil
}
