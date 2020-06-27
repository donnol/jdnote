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
	// 获取请求ID
	RequestID() string

	// 取消
	Cancel()

	// 设置Context
	SetContext(context.Context)

	// 设置用户ID
	SetUserID(userID int)

	// 设置请求ID
	SetRequestID(string)

	// 返回一个新的Context，并设置tx
	NewWithTx(db.DB) Context
}

// myContext myContext
type myContext struct {
	context.Context
	db        db.DB
	logger    utillog.Logger
	userID    int
	requestID string

	cancel context.CancelFunc
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

// RequestID 获取请求ID
func (mc *myContext) RequestID() string {
	return ""
}

// Cancel 取消
func (mc *myContext) Cancel() {
	mc.cancel()
}

// SetContext 设置Context
func (mc *myContext) SetContext(ctx context.Context) {
	mc.Context = ctx
}

// SetUserID 设置用户ID
func (mc *myContext) SetUserID(userID int) {
	mc.userID = userID
}

func (mc *myContext) SetRequestID(reqID string) {
	mc.requestID = reqID
}

// NewWithTx 返回一个新的Context，并设置tx
func (mc *myContext) NewWithTx(tx db.DB) Context {
	return New(tx, mc.logger, mc.userID)
}

// New 新建
func New(db db.DB, logger utillog.Logger, userID int) Context {
	mctx := new(myContext)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	mctx.Context = ctx
	mctx.cancel = cancel
	mctx.db = db
	mctx.logger = logger
	mctx.userID = userID
	return mctx
}
