package context

import (
	"context"

	"github.com/donnol/jdnote/utils/store/db"
	utillog "github.com/donnol/tools/log"
	"github.com/pkg/errors"
)

const (
	UserKey    = "UserID"
	RequestKey = "RequestID"
)

// Context 上下文
type Context interface {
	context.Context

	// 获取DB实例
	DB() db.DB
	// 获取日志实例
	Logger() utillog.Logger

	// 取消
	Cancel()

	// 设置Context
	SetContext(context.Context)
	// 返回一个新的Context，并设置tx
	NewWithTx(db.DB) Context
}

// myContext myContext
type myContext struct {
	context.Context

	db     db.DB
	logger utillog.Logger

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

// Cancel 取消
func (mc *myContext) Cancel() {
	mc.cancel()
}

// SetContext 设置Context
func (mc *myContext) SetContext(ctx context.Context) {
	mc.Context = ctx
}

// NewWithTx 返回一个新的Context，并设置tx
func (mc *myContext) NewWithTx(tx db.DB) Context {
	mctx := newCtx(mc.Context, mc.cancel, tx, mc.logger)
	return mctx
}

// New 新建
func New(db db.DB, logger utillog.Logger, userID int) Context {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	mctx := newCtx(ctx, cancel, db, logger)

	return mctx
}

func newCtx(ctx context.Context, cancel context.CancelFunc, db db.DB, logger utillog.Logger) Context {
	mctx := new(myContext)

	mctx.Context = ctx
	mctx.cancel = cancel

	mctx.db = db
	mctx.logger = logger

	return mctx
}

func WithValue(ctx Context, key, value interface{}) Context {
	nctx := context.WithValue(ctx, key, value)
	return newCtx(nctx, ctx.Cancel, ctx.DB(), ctx.Logger())
}

func GetValue(ctx Context, key interface{}) interface{} {
	return ctx.Value(key)
}

func GetUserValue(ctx Context) (int, error) {
	v := GetValue(ctx, UserKey)
	vv, ok := v.(int)
	if !ok {
		return 0, errors.Errorf("get %s failed, got %v", UserKey, v)
	}
	return vv, nil
}

func GetRequestValue(ctx Context) (string, error) {
	v := GetValue(ctx, RequestKey)
	vv, ok := v.(string)
	if !ok {
		return "", errors.Errorf("get %s failed, got %v", RequestKey, v)
	}
	return vv, nil
}

func MustGetUserValue(ctx Context) int {
	v, err := GetUserValue(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func MustGetRequestValue(ctx Context) string {
	v, err := GetRequestValue(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
