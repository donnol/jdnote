package authsrv

import (
	"context"

	"github.com/donnol/jdnote/stores/roleactionstore"
	"github.com/donnol/jdnote/stores/userstore"
)

type IAuth interface {
	CheckPerm(ctx context.Context, perms []string) error
	CheckUserExist(ctx context.Context) error
	CheckLogin(ctx context.Context) error
	CheckUserPerm(ctx context.Context, perms []string) error
}

func New(
	RoleActionstore roleactionstore.IRoleAction,
	Userstore userstore.IUser,
) IAuth {
	return &authImpl{
		RoleActionStore: RoleActionstore,
		UserStore:       Userstore,
	}
}
