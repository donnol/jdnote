package context

import (
	"context"

	utillog "github.com/donnol/jdnote/utils/log"
	"github.com/donnol/jdnote/utils/store/db"
)

// Context 上下文
type Context interface {
	context.Context

	// 获取DB实例
	DB() db.DB
	// 获取日志实例
	Logger() utillog.Logger
	// 获取当前登录用户ID
	UserID() int
}

// myContext myContext
type myContext struct {
	context.Context
	db     db.DB
	logger utillog.Logger
	userID int
}

// DB 获取DB实例
func (mc *myContext) DB() (db db.DB) {
	return mc.db
}

// Logger 获取日志实例
func (mc *myContext) Logger() utillog.Logger {
	return mc.logger
}

// UserID 获取当前登录用户ID
func (mc *myContext) UserID() int {
	return mc.userID
}

// New 新建
func New(db db.DB, logger utillog.Logger, userID int) Context {
	mctx := new(myContext)
	mctx.Context = context.Background()
	mctx.db = db
	mctx.logger = logger
	mctx.userID = userID
	return mctx
}
