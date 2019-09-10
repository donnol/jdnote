package auth

import (
	"github.com/donnol/jdnote/models/roleaction/roleactiondb"
	"github.com/donnol/jdnote/utils/context"
)

// Auth 认证
type Auth struct {
	RoleActionModel roleactiondb.RoleAction
}

// CheckPerm 检查用户是否拥有指定权限
func (a *Auth) CheckPerm(ctx context.Context, perms []string) error {
	if err := a.RoleActionModel.CheckPerm(ctx, perms); err != nil {
		return err
	}

	return nil
}