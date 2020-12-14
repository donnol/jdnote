package userstore

import (
	"github.com/donnol/jdnote/models/usermodel"
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/tools/inject"
)

type UserMock struct {
	AddFunc func(ctx context.Context, e usermodel.Entity) (id int, err error)

	GetByIDFunc func(ctx context.Context, id int) (e usermodel.Entity, err error)

	GetByNameFunc func(ctx context.Context, name string) (e usermodel.Entity, err error)

	GetFirstFunc func(ctx context.Context) (e usermodel.Entity, err error)

	VerifyByNameAndPasswordFunc func(ctx context.Context, name string, password string) (e usermodel.Entity, err error)
}

var (
	_ IUser = &UserMock{}

	userMockCommonProxyContext = inject.ProxyContext{
		PkgPath:       "github.com/donnol/jdnote/stores/userstore",
		InterfaceName: "IUser",
	}
	UserMockAddProxyContext = func() (pctx inject.ProxyContext) {
		pctx = userMockCommonProxyContext
		pctx.MethodName = "Add"
		return
	}()
	UserMockGetByIDProxyContext = func() (pctx inject.ProxyContext) {
		pctx = userMockCommonProxyContext
		pctx.MethodName = "GetByID"
		return
	}()
	UserMockGetByNameProxyContext = func() (pctx inject.ProxyContext) {
		pctx = userMockCommonProxyContext
		pctx.MethodName = "GetByName"
		return
	}()
	UserMockGetFirstProxyContext = func() (pctx inject.ProxyContext) {
		pctx = userMockCommonProxyContext
		pctx.MethodName = "GetFirst"
		return
	}()
	UserMockVerifyByNameAndPasswordProxyContext = func() (pctx inject.ProxyContext) {
		pctx = userMockCommonProxyContext
		pctx.MethodName = "VerifyByNameAndPassword"
		return
	}()
)

func (mockRecv *UserMock) Add(ctx context.Context, e usermodel.Entity) (id int, err error) {
	return mockRecv.AddFunc(ctx, e)
}

func (mockRecv *UserMock) GetByID(ctx context.Context, id int) (e usermodel.Entity, err error) {
	return mockRecv.GetByIDFunc(ctx, id)
}

func (mockRecv *UserMock) GetByName(ctx context.Context, name string) (e usermodel.Entity, err error) {
	return mockRecv.GetByNameFunc(ctx, name)
}

func (mockRecv *UserMock) GetFirst(ctx context.Context) (e usermodel.Entity, err error) {
	return mockRecv.GetFirstFunc(ctx)
}

func (mockRecv *UserMock) VerifyByNameAndPassword(ctx context.Context, name string, password string) (e usermodel.Entity, err error) {
	return mockRecv.VerifyByNameAndPasswordFunc(ctx, name, password)
}

type EntityMock struct {
	FilterFunc func() interface{}
}

var (
	_ IEntity = &EntityMock{}

	entityMockCommonProxyContext = inject.ProxyContext{
		PkgPath:       "github.com/donnol/jdnote/stores/userstore",
		InterfaceName: "IEntity",
	}
	EntityMockFilterProxyContext = func() (pctx inject.ProxyContext) {
		pctx = entityMockCommonProxyContext
		pctx.MethodName = "Filter"
		return
	}()
)

func (mockRecv *EntityMock) Filter() interface{} {
	return mockRecv.FilterFunc()
}
