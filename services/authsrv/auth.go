package authsrv

import (
	"context"

	"github.com/donnol/jdnote/stores/roleactionstore"
	"github.com/donnol/jdnote/stores/userstore"
	utilctx "github.com/donnol/tools/context"
	"github.com/donnol/tools/errors"
)

// authImpl 认证
type authImpl struct {
	RoleActionStore roleactionstore.IRoleAction
	UserStore       userstore.IUser
}

// CheckUserExist 检查用户是否存在
func (a *authImpl) CheckUserExist(ctx context.Context) error {
	userID, err := utilctx.GetUserValue(ctx)
	if err != nil {
		return err
	}
	_, err = a.UserStore.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

// CheckPerm 检查用户是否拥有指定权限
func (a *authImpl) CheckPerm(ctx context.Context, perms []string) error {
	if err := a.RoleActionStore.CheckPerm(ctx, perms); err != nil {
		return err
	}

	return nil
}

// CheckLogin 检查登录态
func (a *authImpl) CheckLogin(ctx context.Context) error {
	userID, err := utilctx.GetUserValue(ctx)
	if err != nil {
		return err
	}
	if userID == 0 {
		return errors.NewNormal(errors.ErrorCodeAuth, "please login")
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
