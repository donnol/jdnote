package log

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"
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

var (
	logRand *rand.Rand
)

func init() {
	// 初始化随机数生成器
	source := rand.NewSource(time.Now().Unix())
	logRand = rand.New(source)
}

// Logger 日志
type Logger struct {
	*log.Logger
}

// New 新建
func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{
		Logger: log.New(out, prefix+" ", flag),
	}
}

// Fatalf 致命
func (l *Logger) Fatalf(format string, v ...interface{}) {
	format = getFormat(FatalLevel, format)
	l.printf(format, v...)
}

// Errorf 错误
func (l *Logger) Errorf(format string, v ...interface{}) {
	format = getFormat(ErrorLevel, format)
	l.printf(format, v...)
}

// Warnf 警告
func (l *Logger) Warnf(format string, v ...interface{}) {
	format = getFormat(WarnLevel, format)
	l.printf(format, v...)
}

// Infof 信息
func (l *Logger) Infof(format string, v ...interface{}) {
	format = getFormat(InfoLevel, format)
	l.printf(format, v...)
}

// Debugf 调试
func (l *Logger) Debugf(format string, v ...interface{}) {
	format = getFormat(DebugLevel, format)
	l.printf(format, v...)
}

// Tracef 追踪
func (l *Logger) Tracef(format string, v ...interface{}) {
	format = getFormat(TraceLevel, format)
	l.printf(format, v...)
}

func (l *Logger) printf(format string, v ...interface{}) {
	l.Logger.Output(3, fmt.Sprintf(format, v...))
}

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
	r := logRand.Int63()
	return fmt.Sprintf("|%d| [%s] %s", r, level, format)
}

func printf(format string, v ...interface{}) {
	log.Output(3, fmt.Sprintf(format, v...))
}
