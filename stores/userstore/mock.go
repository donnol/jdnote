package userstore

import (
	"github.com/donnol/jdnote/models/usermodel"
	"github.com/donnol/jdnote/utils/context"
)

type UserMock struct {
	AddFunc func(ctx context.Context, e usermodel.Entity) (id int, err error)

	GetByIDFunc func(ctx context.Context, id int) (e usermodel.Entity, err error)

	GetByNameFunc func(ctx context.Context, name string) (e usermodel.Entity, err error)

	GetFirstFunc func(ctx context.Context) (e usermodel.Entity, err error)

	VerifyByNameAndPasswordFunc func(ctx context.Context, name string, password string) (e usermodel.Entity, err error)
}

var _ IUser = &UserMock{}

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

var _ IEntity = &EntityMock{}

func (mockRecv *EntityMock) Filter() interface{} {
	return mockRecv.FilterFunc()
}
