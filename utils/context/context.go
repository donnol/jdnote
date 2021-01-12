package context

import (
	"context"

	"github.com/donnol/jdnote/utils/store/db"
)

// Context 上下文
type Context interface {
	context.Context

	// 获取DB实例
	DB() db.DB

	// 返回一个新的Context，并设置tx
	NewWithTx(db.DB) Context

	SetContext(ctx context.Context)
	StdContext() context.Context
}

// myContext myContext
type myContext struct {
	context.Context

	db db.DB
}

// DB 获取DB实例
func (mc *myContext) DB() (db db.DB) {
	return mc.db
}

// NewWithTx 返回一个新的Context，并设置tx
func (mc *myContext) NewWithTx(tx db.DB) Context {
	mctx := newCtx(mc.Context, tx)
	return mctx
}

// SetContext 设置Context
func (mc *myContext) SetContext(ctx context.Context) {
	mc.Context = ctx
}

func (mc *myContext) StdContext() context.Context {
	return mc.Context
}

// New 新建
func New(ctx context.Context, db db.DB) Context {
	mctx := newCtx(ctx, db)

	return mctx
}

func newCtx(ctx context.Context, db db.DB) Context {
	mctx := new(myContext)

	mctx.Context = ctx

	mctx.db = db

	return mctx
}

// WithValue 往标准库ctx设置key,value
func WithValue(ctx Context, key, value interface{}) Context {
	nctx := context.WithValue(ctx.StdContext(), key, value)
	return newCtx(nctx, ctx.DB())
}

// GetValue 从标准库ctx读取key对应value
func GetValue(ctx Context, key interface{}) interface{} {
	return ctx.StdContext().Value(key)
}
