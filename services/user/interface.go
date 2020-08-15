package user

import (
	"github.com/donnol/jdnote/app"
	"github.com/donnol/jdnote/models/user"
	"github.com/donnol/jdnote/models/userrole"
	"github.com/donnol/jdnote/utils/context"
)

type IUser interface {
	Check(ctx context.Context) error
	GetByID(ctx context.Context, id int) (e user.Entity, err error)
	GetByName(ctx context.Context, name string) (e user.Entity, err error)
	GetFirst(ctx context.Context) (e user.Entity, err error)
	VerifyByNameAndPassword(ctx context.Context, name, password string) (e user.Entity, err error)
	Add(ctx context.Context, e user.Entity) (id int, err error)
}

func New(
	UserModel user.IUser,
	UserRoleModel userrole.IUserRole,
) IUser {
	return &userImpl{
		UserModel:     UserModel,
		UserRoleModel: UserRoleModel,
	}
}

func init() {
	app.MustRegisterProvider(New)
}
