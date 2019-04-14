package log

import (
	"log"
	"os"
	"reflect"

	"github.com/donnol/jdnote/utils/reflectx"
)

// InitParamWithLogger 初始化param里的Logger
func InitParamWithLogger(param interface{}) (interface{}, error) {
	// 类型和值
	specType := reflect.TypeOf((*Logger)(nil)).Elem()
	specValue := reflect.ValueOf(New(os.Stdout, "", log.LstdFlags|log.Llongfile))

	// 注入
	var err error
	param, err = reflectx.InitParam(param, specType, specValue)
	if err != nil {
		return param, err
	}

	return param, nil
}
