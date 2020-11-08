package usermodel

import "github.com/donnol/jdnote/utils/context"

type UserMock struct {
	AddFunc                     func(ctx context.Context, e Entity) (id int, err error)
	GetByIDFunc                 func(ctx context.Context, id int) (e Entity, err error)
	GetByNameFunc               func(ctx context.Context, name string) (e Entity, err error)
	GetFirstFunc                func(ctx context.Context) (e Entity, err error)
	VerifyByNameAndPasswordFunc func(ctx context.Context, name string, password string) (e Entity, err error)
}

func (mock *UserMock) Add(ctx context.Context, e Entity) (id int, err error) {
	return mock.AddFunc(ctx, e)
}

func (mock *UserMock) GetByID(ctx context.Context, id int) (e Entity, err error) {
	return mock.GetByIDFunc(ctx, id)
}

func (mock *UserMock) GetByName(ctx context.Context, name string) (e Entity, err error) {
	return mock.GetByNameFunc(ctx, name)
}

func (mock *UserMock) GetFirst(ctx context.Context) (e Entity, err error) {
	return mock.GetFirstFunc(ctx)
}

func (mock *UserMock) VerifyByNameAndPassword(
	ctx context.Context,
	name string,
	password string,
) (e Entity, err error) {
	return mock.VerifyByNameAndPasswordFunc(ctx, name, password)
}
