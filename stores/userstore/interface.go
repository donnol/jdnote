package userstore

import (
	"context"

	"github.com/donnol/jdnote/models/usermodel"
	"github.com/donnol/jdnote/utils/store/db"
)

type IUser interface {
	Add(ctx context.Context, e usermodel.Entity) (id int, err error)
	GetByID(ctx context.Context, id int) (e usermodel.Entity, err error)
	GetByName(ctx context.Context, name string) (e usermodel.Entity, err error)
	GetFirst(ctx context.Context) (e usermodel.Entity, err error)
	VerifyByNameAndPassword(ctx context.Context, name string, password string) (e usermodel.Entity, err error)
}

type IEntity interface{ Filter() interface{} }

func New(
	db db.DB,
) IUser {
	return &userImpl{
		db: db,
	}
}
