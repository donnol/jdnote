package auth

import (
	"github.com/donnol/jdnote/context"
	roleaction "github.com/donnol/jdnote/model/role_action"
)

// Auth 认证
type Auth struct {
	RoleActionModel roleaction.RoleAction
}

// CheckPerm 检查用户是否拥有指定权限
func (a *Auth) CheckPerm(ctx context.Context, perms []string) error {
	if err := a.RoleActionModel.CheckPerm(ctx, perms); err != nil {
		return err
	}

	return nil
}
