package auth

import (
	"github.com/donnol/jdnote/app"
	"github.com/donnol/jdnote/models/roleaction"
	"github.com/donnol/jdnote/models/user"
	"github.com/donnol/jdnote/utils/context"
)

type IAuth interface {
	CheckPerm(ctx context.Context, perms []string) error
	CheckUserExist(ctx context.Context) error
	CheckLogin(ctx context.Context) error
	CheckUserPerm(ctx context.Context, perms []string) error
}

func New(
	RoleActionModel roleaction.IRoleAction,
	UserModel user.IUser,
) IAuth {
	return &authImpl{
		RoleActionModel: RoleActionModel,
		UserModel:       UserModel,
	}
}

func init() {
	app.MustRegisterProvider(New)
}
