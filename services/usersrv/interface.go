package usersrv

import (
	"context"

	"github.com/donnol/jdnote/models/usermodel"
	"github.com/donnol/jdnote/stores/userrolestore"
	"github.com/donnol/jdnote/stores/userstore"
)

type IUser interface {
	Check(ctx context.Context) error
	GetByID(ctx context.Context, id int) (e usermodel.Entity, err error)
	GetByName(ctx context.Context, name string) (e usermodel.Entity, err error)
	GetFirst(ctx context.Context) (e usermodel.Entity, err error)
	VerifyByNameAndPassword(ctx context.Context, name, password string) (e usermodel.Entity, err error)
	Add(ctx context.Context, e usermodel.Entity) (id int, err error)
}

func New(
	Userstore userstore.IUser,
	UserRolestore userrolestore.IUserRole,
) IUser {
	return &userImpl{
		UserStore:     Userstore,
		UserRoleStore: UserRolestore,
	}
}
