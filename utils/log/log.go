package log

import (
	"fmt"
	"log"
)

// 日志级别--https://stackoverflow.com/questions/2031163/when-to-use-the-different-log-levels
const (
	FatalLevel = "FATAL"
	ErrorLevel = "ERROR"
	WarnLevel  = "WARN"
	InfoLevel  = "INFO"
	DebugLevel = "DEBUG"
	TraceLevel = "TRACE"
)

// Fatalf 致命
func Fatalf(format string, v ...interface{}) {
	format = getFormat(FatalLevel, format)
	printf(format, v...)
}

// Errorf 错误
func Errorf(format string, v ...interface{}) {
	format = getFormat(ErrorLevel, format)
	printf(format, v...)
}

// Warnf 警告
func Warnf(format string, v ...interface{}) {
	format = getFormat(WarnLevel, format)
	printf(format, v...)
}

// Infof 信息
func Infof(format string, v ...interface{}) {
	format = getFormat(InfoLevel, format)
	printf(format, v...)
}

// Debugf 调试
func Debugf(format string, v ...interface{}) {
	format = getFormat(DebugLevel, format)
	printf(format, v...)
}

// Tracef 追踪
func Tracef(format string, v ...interface{}) {
	format = getFormat(TraceLevel, format)
	printf(format, v...)
}

func getFormat(level, format string) string {
	return fmt.Sprintf("[%s] %s", level, format)
}

func printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
