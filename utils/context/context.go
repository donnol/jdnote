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

	// 设置Context
	SetContext(context.Context)

	// 设置用户ID
	SetUserID(userID int)

	// 返回一个新的Context，并设置tx
	NewWithTx(db.DB) Context
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

// SetContext 设置Context
func (mc *myContext) SetContext(ctx context.Context) {
	mc.Context = ctx
}

// SetUserID 设置用户ID
func (mc *myContext) SetUserID(userID int) {
	mc.userID = userID
}

// NewWithTx 返回一个新的Context，并设置tx
func (mc *myContext) NewWithTx(tx db.DB) Context {
	nmctx := new(myContext)
	nmctx.Context = mc.Context
	nmctx.db = tx
	nmctx.logger = mc.logger
	nmctx.userID = mc.userID
	return nmctx
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
