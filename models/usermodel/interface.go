package usermodel

import (
	"github.com/donnol/jdnote/utils/context"
)

type IUser interface {
	Add(ctx context.Context, e Entity) (id int, err error)
	GetByID(ctx context.Context, id int) (e Entity, err error)
	GetByName(ctx context.Context, name string) (e Entity, err error)
	GetFirst(ctx context.Context) (e Entity, err error)
	VerifyByNameAndPassword(ctx context.Context, name string, password string) (e Entity, err error)
}

type IEntity interface{ Filter() interface{} }

func New() IUser {
	return &userImpl{}
}
