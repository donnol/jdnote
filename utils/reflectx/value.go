package reflectx

import (
	"fmt"
	"reflect"
	"time"
)

var (
	// ErrParamNotStruct 参数不是结构体
	ErrParamNotStruct = fmt.Errorf("Please input struct param")
)

// InitParam 初始化-使用反射初始化param里的指定类型
func InitParam(param interface{}, specType reflect.Type, specValue reflect.Value, copy bool) (interface{}, error) {
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

	// 创建副本
	if copy {
		var sf = make([]reflect.StructField, 0)
		for i := 0; i < refType.NumField(); i++ {
			field := refType.Field(i)

			sf = append(sf, field)
		}
		newType := reflect.StructOf(sf)
		newValue := reflect.New(refType)

		// 给value赋值
		newValueElem := newValue.Elem()
		for i := 0; i < refType.NumField(); i++ {
			oldV := refValue.Field(i)
			newV := newValueElem.Field(i)
			newV.Set(oldV)
		}

		// 替换
		refType = newType
		refValue = newValue
	}

	// 注入value
	setValue(refType, specType, refValue, specValue)

	// 返回副本
	if copy {
		return refValue.Interface(), nil
	}

	return param, nil
}

func setValue(refType, specType reflect.Type, refValue, specValue reflect.Value) {
	// 忽略非结构体或者time.Time类型
	if refType.Kind() != reflect.Struct ||
		refType == reflect.TypeOf((*time.Time)(nil)).Elem() {
		return
	}

	for i := 0; i < refType.NumField(); i++ {
		field := refType.Field(i)

		// 忽略非导出字段
		if field.PkgPath != "" {
			continue
		}

		// 获取对应字段的value
		var value reflect.Value
		if refValue.Type().Kind() == reflect.Ptr {
			value = refValue.Elem().Field(i)
		} else {
			value = refValue.Field(i)
		}

		// 根据字段type判断是否可以赋值
		if field.Type == specType { // 类型相同，直接赋值
			value.Set(specValue)
		} else { // 匿名内嵌或者包含在普通字段里，继续对该字段类型遍历
			setValue(field.Type, specType, value, specValue)
		}
	}
}
