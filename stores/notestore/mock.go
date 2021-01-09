package notestore

import (
	"github.com/donnol/jdnote/models/notemodel"
	"github.com/donnol/jdnote/utils/common"
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/tools/inject"
)

type NoterMock struct {
	AddFunc func(ctx context.Context, entity notemodel.Entity) (id int, err error)

	AddOneFunc func(ctx context.Context) (id int, err error)

	DelFunc func(ctx context.Context, id int) (err error)

	GetFunc func(ctx context.Context, id int) (entity notemodel.Entity, err error)

	GetListFunc func(ctx context.Context, ids []int64) (entitys []notemodel.Entity, err error)

	GetPageFunc func(ctx context.Context, entity notemodel.Entity, param common.Param) (res []notemodel.EntityWithTotal, err error)

	ModFunc func(ctx context.Context, id int, entity *notemodel.Entity) (err error)

	ModStatusFunc func(ctx context.Context, id int, status notemodel.Status) (err error)
}

var (
	_ Noter = &NoterMock{}

	noterMockCommonProxyContext = inject.ProxyContext{
		PkgPath:       "github.com/donnol/jdnote/stores/notestore",
		InterfaceName: "Noter",
	}
	NoterMockAddProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noterMockCommonProxyContext
		pctx.MethodName = "Add"
		return
	}()
	NoterMockAddOneProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noterMockCommonProxyContext
		pctx.MethodName = "AddOne"
		return
	}()
	NoterMockDelProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noterMockCommonProxyContext
		pctx.MethodName = "Del"
		return
	}()
	NoterMockGetProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noterMockCommonProxyContext
		pctx.MethodName = "Get"
		return
	}()
	NoterMockGetListProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noterMockCommonProxyContext
		pctx.MethodName = "GetList"
		return
	}()
	NoterMockGetPageProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noterMockCommonProxyContext
		pctx.MethodName = "GetPage"
		return
	}()
	NoterMockModProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noterMockCommonProxyContext
		pctx.MethodName = "Mod"
		return
	}()
	NoterMockModStatusProxyContext = func() (pctx inject.ProxyContext) {
		pctx = noterMockCommonProxyContext
		pctx.MethodName = "ModStatus"
		return
	}()
)

func (mockRecv *NoterMock) Add(ctx context.Context, entity notemodel.Entity) (id int, err error) {
	return mockRecv.AddFunc(ctx, entity)
}

func (mockRecv *NoterMock) AddOne(ctx context.Context) (id int, err error) {
	return mockRecv.AddOneFunc(ctx)
}

func (mockRecv *NoterMock) Del(ctx context.Context, id int) (err error) {
	return mockRecv.DelFunc(ctx, id)
}

func (mockRecv *NoterMock) Get(ctx context.Context, id int) (entity notemodel.Entity, err error) {
	return mockRecv.GetFunc(ctx, id)
}

func (mockRecv *NoterMock) GetList(ctx context.Context, ids []int64) (entitys []notemodel.Entity, err error) {
	return mockRecv.GetListFunc(ctx, ids)
}

func (mockRecv *NoterMock) GetPage(ctx context.Context, entity notemodel.Entity, param common.Param) (res []notemodel.EntityWithTotal, err error) {
	return mockRecv.GetPageFunc(ctx, entity, param)
}

func (mockRecv *NoterMock) Mod(ctx context.Context, id int, entity *notemodel.Entity) (err error) {
	return mockRecv.ModFunc(ctx, id, entity)
}

func (mockRecv *NoterMock) ModStatus(ctx context.Context, id int, status notemodel.Status) (err error) {
	return mockRecv.ModStatusFunc(ctx, id, status)
}
