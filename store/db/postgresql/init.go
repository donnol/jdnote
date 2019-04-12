package pg

import (
	"reflect"

	"github.com/donnol/jdnote/utils/reflectx"
)

// InitParamWithDB 初始化param里的DB
func InitParamWithDB(param interface{}) (interface{}, error) {
	// db类型和值
	dbType := reflect.TypeOf((*DB)(nil)).Elem()
	dbValue := reflect.ValueOf(New())

	// 注入
	var err error
	param, err = reflectx.InitParam(param, dbType, dbValue)
	if err != nil {
		return param, err
	}

	return param, nil
}
