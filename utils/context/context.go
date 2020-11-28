package context

import (
	"context"

	"github.com/donnol/jdnote/utils/store/db"
	"github.com/pkg/errors"
)

type (
	TimestampType  string
	RemoteAddrType string
	UserKeyType    string
	RequestKeyType string
)

const (
	// 时间
	TimestampKey TimestampType = "Timestamp"

	// 地点
	RemoteAddrKey RemoteAddrType = "RemoteAddr"

	// 用户
	UserKey UserKeyType = "UserID"

	// 请求
	RequestKey RequestKeyType = "RequestID"
)

// Context 上下文
type Context interface {
	context.Context

	// 获取DB实例
	DB() db.DB

	// 返回一个新的Context，并设置tx
	NewWithTx(db.DB) Context
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

// SetContext 设置Context
func (mc *myContext) SetContext(ctx context.Context) {
	mc.Context = ctx
}

// NewWithTx 返回一个新的Context，并设置tx
func (mc *myContext) NewWithTx(tx db.DB) Context {
	mctx := newCtx(mc.Context, tx)
	return mctx
}

// New 新建
func New(ctx context.Context, db db.DB, userID int) Context {
	ctx = context.WithValue(ctx, UserKey, userID)
	mctx := newCtx(ctx, db)

	return mctx
}

func newCtx(ctx context.Context, db db.DB) Context {
	mctx := new(myContext)

	mctx.Context = ctx

	mctx.db = db

	return mctx
}

func WithValue(ctx Context, key, value interface{}) Context {
	nctx := context.WithValue(ctx, key, value)
	return newCtx(nctx, ctx.DB())
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

func GetTimestampValue(ctx Context) (int64, error) {
	v := GetValue(ctx, TimestampKey)
	vv, ok := v.(int64)
	if !ok {
		return 0, errors.Errorf("get %s failed, got %v", TimestampKey, v)
	}
	return vv, nil
}

func GetRemoteAddrValue(ctx Context) (string, error) {
	v := GetValue(ctx, RemoteAddrKey)
	vv, ok := v.(string)
	if !ok {
		return "", errors.Errorf("get %s failed, got %v", RemoteAddrKey, v)
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

func MustGetTimestampValue(ctx Context) int64 {
	v, err := GetTimestampValue(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func MustGetRemoteAddrValue(ctx Context) string {
	v, err := GetRemoteAddrValue(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
