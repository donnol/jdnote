package context

import (
	"context"

	pg "github.com/donnol/jdnote/store/db/postgresql"
	utillog "github.com/donnol/jdnote/utils/log"
)

// Context 上下文
type Context interface {
	context.Context

	// 获取DB实例
	DB() pg.DB
	// 获取日志实例
	Logger() utillog.Logger
}

// myContext myContext
type myContext struct {
	context.Context
	db     pg.DB
	logger utillog.Logger
}

// DB 获取DB实例
func (mc *myContext) DB() (db pg.DB) {
	return mc.db
}

// Logger 获取日志实例
func (mc *myContext) Logger() utillog.Logger {
	return mc.logger
}

// New 新建
func New(db pg.DB, logger utillog.Logger) Context {
	mctx := new(myContext)
	mctx.Context = context.Background()
	mctx.db = db
	mctx.logger = logger
	return mctx
}
