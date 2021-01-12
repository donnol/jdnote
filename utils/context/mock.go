package context

import (
	"context"
	"time"

	"github.com/donnol/jdnote/utils/store/db"
	"github.com/donnol/tools/inject"
)

type ContextMock struct {
	DBFunc func() db.DB

	DeadlineFunc func() (deadline time.Time, ok bool)

	DoneFunc func() <-chan struct{}

	ErrFunc func() error

	NewWithTxFunc func(db.DB) Context

	SetContextFunc func(ctx context.Context)

	StdContextFunc func() context.Context

	ValueFunc func(key interface{}) interface{}
}

var (
	_ Context = &ContextMock{}

	contextMockCommonProxyContext = inject.ProxyContext{
		PkgPath:       "github.com/donnol/jdnote/utils/context",
		InterfaceName: "Context",
	}
	ContextMockDBProxyContext = func() (pctx inject.ProxyContext) {
		pctx = contextMockCommonProxyContext
		pctx.MethodName = "DB"
		return
	}()
	ContextMockDeadlineProxyContext = func() (pctx inject.ProxyContext) {
		pctx = contextMockCommonProxyContext
		pctx.MethodName = "Deadline"
		return
	}()
	ContextMockDoneProxyContext = func() (pctx inject.ProxyContext) {
		pctx = contextMockCommonProxyContext
		pctx.MethodName = "Done"
		return
	}()
	ContextMockErrProxyContext = func() (pctx inject.ProxyContext) {
		pctx = contextMockCommonProxyContext
		pctx.MethodName = "Err"
		return
	}()
	ContextMockNewWithTxProxyContext = func() (pctx inject.ProxyContext) {
		pctx = contextMockCommonProxyContext
		pctx.MethodName = "NewWithTx"
		return
	}()
	ContextMockSetContextProxyContext = func() (pctx inject.ProxyContext) {
		pctx = contextMockCommonProxyContext
		pctx.MethodName = "SetContext"
		return
	}()
	ContextMockStdContextProxyContext = func() (pctx inject.ProxyContext) {
		pctx = contextMockCommonProxyContext
		pctx.MethodName = "StdContext"
		return
	}()
	ContextMockValueProxyContext = func() (pctx inject.ProxyContext) {
		pctx = contextMockCommonProxyContext
		pctx.MethodName = "Value"
		return
	}()
)

func (mockRecv *ContextMock) DB() db.DB {
	return mockRecv.DBFunc()
}

func (mockRecv *ContextMock) Deadline() (deadline time.Time, ok bool) {
	return mockRecv.DeadlineFunc()
}

func (mockRecv *ContextMock) Done() <-chan struct{} {
	return mockRecv.DoneFunc()
}

func (mockRecv *ContextMock) Err() error {
	return mockRecv.ErrFunc()
}

func (mockRecv *ContextMock) NewWithTx(p0 db.DB) Context {
	return mockRecv.NewWithTxFunc(p0)
}

func (mockRecv *ContextMock) SetContext(ctx context.Context) {
	mockRecv.SetContextFunc(ctx)
}

func (mockRecv *ContextMock) StdContext() context.Context {
	return mockRecv.StdContextFunc()
}

func (mockRecv *ContextMock) Value(key interface{}) interface{} {
	return mockRecv.ValueFunc(key)
}
