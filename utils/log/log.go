package log

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"runtime"
	"time"
)

const (
	calldepth = 3
	skip      = 5
)

var (
	logRand *rand.Rand
)

func init() {
	// 初始化随机数生成器
	source := rand.NewSource(time.Now().Unix())
	logRand = rand.New(source)
}

// Notifier 通知接口
type Notifier interface {
	Levels() []Level
	Notify(msg string)
}

// Logger 日志
type Logger struct {
	*log.Logger

	notify Notifier
}

// New 新建
func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{
		Logger: log.New(out, prefix+" ", flag),
	}
}

// SetNotify 设置通知
func (l *Logger) SetNotify(notify Notifier) {
	l.notify = notify
}

// Fatalf 致命
func (l *Logger) Fatalf(format string, v ...interface{}) {
	level := FatalLevel
	l.printf(level, format, v...)
}

// Errorf 错误
func (l *Logger) Errorf(format string, v ...interface{}) {
	level := ErrorLevel
	l.printf(level, format, v...)
}

// Warnf 警告
func (l *Logger) Warnf(format string, v ...interface{}) {
	level := WarnLevel
	l.printf(level, format, v...)
}

// Infof 信息
func (l *Logger) Infof(format string, v ...interface{}) {
	level := InfoLevel
	l.printf(level, format, v...)
}

// Debugf 调试
func (l *Logger) Debugf(format string, v ...interface{}) {
	level := DebugLevel
	l.printf(level, format, v...)
}

// Tracef 追踪
func (l *Logger) Tracef(format string, v ...interface{}) {
	level := TraceLevel
	l.printf(level, format, v...)
}

func (l *Logger) printf(level Level, format string, v ...interface{}) {
	format = getFormat(level, format)
	msg := fmt.Sprintf(format, v...)

	l.Logger.Output(calldepth, msg)

	// 发送通知
	l.notice(level, msg)
}

// 发送通知
func (l *Logger) notice(level Level, msg string) {
	if l.notify == nil {
		return
	}

	levels := l.notify.Levels()
	if !InLevel(levels, level) {
		return
	}

	stack := collectStack()
	l.notify.Notify(msg + stack)
}

// Fatalf 致命
func Fatalf(format string, v ...interface{}) {
	level := FatalLevel
	printf(level, format, v...)
}

// Errorf 错误
func Errorf(format string, v ...interface{}) {
	level := ErrorLevel
	printf(level, format, v...)
}

// Warnf 警告
func Warnf(format string, v ...interface{}) {
	level := WarnLevel
	printf(level, format, v...)
}

// Infof 信息
func Infof(format string, v ...interface{}) {
	level := InfoLevel
	printf(level, format, v...)
}

// Debugf 调试
func Debugf(format string, v ...interface{}) {
	level := DebugLevel
	printf(level, format, v...)
}

// Tracef 追踪
func Tracef(format string, v ...interface{}) {
	level := TraceLevel
	printf(level, format, v...)
}

func getFormat(level Level, format string) string {
	r := logRand.Int63()
	return fmt.Sprintf("|%d| [%s] %s", r, level, format)
}

func printf(level Level, format string, v ...interface{}) {
	format = getFormat(level, format)
	msg := fmt.Sprintf(format, v...)

	log.Output(calldepth, msg)
}

func collectStack() string {
	var stack string

	// 收集栈信息
	var pcs = make([]uintptr, 10)
	n := runtime.Callers(skip, pcs)
	if n == 0 {
		return stack
	}
	pcs = pcs[:n]
	frames := runtime.CallersFrames(pcs)
	for {
		next, more := frames.Next()
		if !more {
			break
		}
		stack += fmt.Sprintf("%v:%v\n%v\n", next.File, next.Line, next.Function)
	}

	return fmt.Sprintf("\n%s\n", stack)
}
