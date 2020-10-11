package usersrv

import (
	"github.com/donnol/jdnote/models/usermodel"
	"github.com/donnol/jdnote/models/userrolemodel"
	"github.com/donnol/jdnote/utils/context"
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
	UserModel usermodel.IUser,
	UserRoleModel userrolemodel.IUserRole,
) IUser {
	return &userImpl{
		UserModel:     UserModel,
		UserRoleModel: UserRoleModel,
	}
}
