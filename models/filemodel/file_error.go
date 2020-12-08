package filemodel

import "errors"

// 参考：https://mp.weixin.qq.com/s/DPtujfVNMw_Jgel_zGTFvw
var (
	_ error = fileError("")
)

type fileError string

func (fe fileError) Error() string {
	return string(fe)
}

// 错误定义
const (
	// ErrNotFound error = fileError("Not Found") // Compiler: invalid constant type error

	ErrNotFound = fileError("Not Found")
)

func IsErrNotFound(err error) bool {
	// return err == ErrNotFound

	return errors.Is(err, ErrNotFound)
}
