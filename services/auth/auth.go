package auth

import (
	"github.com/donnol/jdnote/models/roleaction"
	"github.com/donnol/jdnote/models/user"
	"github.com/donnol/jdnote/utils/context"
	"github.com/pkg/errors"
)

// authImpl 认证
type authImpl struct {
	RoleActionModel roleaction.IRoleAction
	UserModel       user.IUser
}

// CheckUserExist 检查用户是否存在
func (a *authImpl) CheckUserExist(ctx context.Context) error {
	userID, err := context.GetUserValue(ctx)
	if err != nil {
		return err
	}
	_, err = a.UserModel.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

// CheckPerm 检查用户是否拥有指定权限
func (a *authImpl) CheckPerm(ctx context.Context, perms []string) error {
	if err := a.RoleActionModel.CheckPerm(ctx, perms); err != nil {
		return err
	}

	return nil
}

// CheckLogin 检查登录态
func (a *authImpl) CheckLogin(ctx context.Context) error {
	userID, err := context.GetUserValue(ctx)
	if err != nil {
		return err
	}
	if userID == 0 {
		return errors.Errorf("Please login")
	}
	err = a.CheckUserExist(ctx)
	if err != nil {
		return err
	}
	return nil
}

// CheckUserPerm 检查用户权限
func (a *authImpl) CheckUserPerm(ctx context.Context, perms []string) error {
	// 先要登录
	if err := a.CheckLogin(ctx); err != nil {
		return err
	}

	// 检查权限
	if err := a.CheckPerm(ctx, perms); err != nil {
		return err
	}

	return nil
}
