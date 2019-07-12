package errors

import "fmt"

// 错误级别
const (
	LevelNormal = 1
	LevelFatal  = 2
)

// Error 错误
type Error struct {
	Code int    `json:"code"` // 请求返回码，一般0表示正常，非0表示异常
	Msg  string `json:"msg"`  // 信息，一般是出错时的描述信息

	level int // 级别
}

// New 新建普通错误
func New(code int, msg string) error {
	return Error{
		Code:  code,
		Msg:   msg,
		level: LevelNormal,
	}
}

// Fatal 新建严重错误
func Fatal(code int, msg string) error {
	return Error{
		Code:  code,
		Msg:   msg,
		level: LevelFatal,
	}
}

// Error 实现error接口
func (e Error) Error() string {
	return fmt.Sprintf("[%s] Code: %d, Msg: %s", e.nameByLevel(), e.Code, e.Msg)
}

// IsNormal 是否普通错误
func (e Error) IsNormal() bool {
	return e.level == LevelNormal
}

// IsFatal 是否严重错误
func (e Error) IsFatal() bool {
	return e.level == LevelFatal
}

func (e Error) nameByLevel() string {
	switch e.level {
	case LevelNormal:
		return "Normal"
	case LevelFatal:
		return "Fatal"
	}
	return ""
}

var _ error = Error{}