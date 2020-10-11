package authsrv

import (
	"github.com/donnol/jdnote/models/roleactionmodel"
	"github.com/donnol/jdnote/models/usermodel"
	"github.com/donnol/jdnote/utils/context"
)

type IAuth interface {
	CheckPerm(ctx context.Context, perms []string) error
	CheckUserExist(ctx context.Context) error
	CheckLogin(ctx context.Context) error
	CheckUserPerm(ctx context.Context, perms []string) error
}

func New(
	RoleActionModel roleactionmodel.IRoleAction,
	UserModel usermodel.IUser,
) IAuth {
	return &authImpl{
		RoleActionModel: RoleActionModel,
		UserModel:       UserModel,
	}
}
