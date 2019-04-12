package reflectx

import (
	"fmt"
	"reflect"
)

var (
	// ErrParamNotStruct 参数不是结构体
	ErrParamNotStruct = fmt.Errorf("Please input struct param")
)

// InitParam 初始化-使用反射初始化param里的指定类型
func InitParam(param interface{}, specType reflect.Type, specValue reflect.Value) (interface{}, error) {
	// 反射获取type和value
	refType := reflect.TypeOf(param)
	refValue := reflect.ValueOf(param)
	if refType.Kind() == reflect.Ptr {
		refType = refType.Elem()
		refValue = refValue.Elem()
	}
	if refType.Kind() != reflect.Struct {
		return param, ErrParamNotStruct
	}

	// 注入value
	setValue(refType, specType, refValue, specValue)

	return param, nil
}

func setValue(refType, dbType reflect.Type, refValue, dbValue reflect.Value) {
	for i := 0; i < refType.NumField(); i++ {
		field := refType.Field(i)
		if field.Type == dbType { // 类型相同，直接赋值
			v := refValue.Field(i)
			v.Set(dbValue)
		} else if field.Type.Implements(dbType) { // 内嵌类型，递归遍历
			setValue(field.Type, dbType, refValue.Field(i), dbValue)
		}
	}
}
