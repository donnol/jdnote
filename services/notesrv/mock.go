package notesrv

import (
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/timer"
	"github.com/donnol/tools/inject"
)

type NoteMock struct {
	AddOneFunc func(ctx context.Context) (id int, err error)

	DelFunc func(ctx context.Context, id int) (err error)

	GetFunc func(ctx context.Context, id int) (r Result, err error)

	GetPageFunc func(ctx context.Context, param PageParam) (r PageResult, err error)

	GetPublishFunc func(ctx context.Context, id int) (r Result, err error)

	HideFunc func(ctx context.Context, id int) error

	ModFunc func(ctx context.Context, id int, p *Param) (err error)

	PublishFunc func(ctx context.Context, id int) error

	TimerFunc func(ctx context.Context) timer.FuncJob
}

var (
	_ INote = &NoteMock{}

	noteMockCommonProxyContext = inject.ProxyContext{
		PkgPath:       "github.com/donnol/jdnote/services/notesrv",
		InterfaceName: "INote",
	}
	NoteMockAddOneProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noteMockCommonProxyContext
		pctx.MethodName = "AddOne"
		return
	}()
	NoteMockDelProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noteMockCommonProxyContext
		pctx.MethodName = "Del"
		return
	}()
	NoteMockGetProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noteMockCommonProxyContext
		pctx.MethodName = "Get"
		return
	}()
	NoteMockGetPageProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noteMockCommonProxyContext
		pctx.MethodName = "GetPage"
		return
	}()
	NoteMockGetPublishProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noteMockCommonProxyContext
		pctx.MethodName = "GetPublish"
		return
	}()
	NoteMockHideProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noteMockCommonProxyContext
		pctx.MethodName = "Hide"
		return
	}()
	NoteMockModProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noteMockCommonProxyContext
		pctx.MethodName = "Mod"
		return
	}()
	NoteMockPublishProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noteMockCommonProxyContext
		pctx.MethodName = "Publish"
		return
	}()
	NoteMockTimerProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noteMockCommonProxyContext
		pctx.MethodName = "Timer"
		return
	}()
)

func (mockRecv *NoteMock) AddOne(ctx context.Context) (id int, err error) {
	return mockRecv.AddOneFunc(ctx)
}

func (mockRecv *NoteMock) Del(ctx context.Context, id int) (err error) {
	return mockRecv.DelFunc(ctx, id)
}

func (mockRecv *NoteMock) Get(ctx context.Context, id int) (r Result, err error) {
	return mockRecv.GetFunc(ctx, id)
}

func (mockRecv *NoteMock) GetPage(ctx context.Context, param PageParam) (r PageResult, err error) {
	return mockRecv.GetPageFunc(ctx, param)
}

func (mockRecv *NoteMock) GetPublish(ctx context.Context, id int) (r Result, err error) {
	return mockRecv.GetPublishFunc(ctx, id)
}

func (mockRecv *NoteMock) Hide(ctx context.Context, id int) error {
	return mockRecv.HideFunc(ctx, id)
}

func (mockRecv *NoteMock) Mod(ctx context.Context, id int, p *Param) (err error) {
	return mockRecv.ModFunc(ctx, id, p)
}

func (mockRecv *NoteMock) Publish(ctx context.Context, id int) error {
	return mockRecv.PublishFunc(ctx, id)
}

func (mockRecv *NoteMock) Timer(ctx context.Context) timer.FuncJob {
	return mockRecv.TimerFunc(ctx)
}
