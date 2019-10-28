package auth

import (
	"github.com/donnol/jdnote/models"
	"github.com/donnol/jdnote/models/roleaction/roleactiondata"
	"github.com/donnol/jdnote/models/user/userdata"
	"github.com/donnol/jdnote/utils/context"
)

// Auth 认证
type Auth struct {
	models.Base

	RoleActionModel roleactiondata.RoleAction
	UserModel       userdata.User
}

// CheckUserExist 检查用户是否存在
func (a *Auth) CheckUserExist(ctx context.Context) error {
	_, err := a.UserModel.GetByID(ctx, ctx.UserID())
	if err != nil {
		return err
	}

	return nil
}

// CheckPerm 检查用户是否拥有指定权限
func (a *Auth) CheckPerm(ctx context.Context, perms []string) error {
	if err := a.RoleActionModel.CheckPerm(ctx, perms); err != nil {
		return err
	}

	return nil
}
