package userrolestore

import (
	"github.com/donnol/jdnote/models/userrolemodel"
	"github.com/donnol/jdnote/utils/context"
	"github.com/pkg/errors"
)

type userRoleImpl struct {
}

// GetByUserID 获取用户相关的角色
func (ur *userRoleImpl) GetByUserID(ctx context.Context, userID int) (list []userrolemodel.Entity, err error) {
	if err = ctx.DB().SelectContext(ctx.StdContext(), &list, `
		SELECT * FROM t_user_role WHERE user_id = $1
		`, userID); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// Add 添加
func (ur *userRoleImpl) Add(ctx context.Context, e userrolemodel.Entity) (id int, err error) {
	if err = ctx.DB().GetContext(ctx.StdContext(), &id, `
		INSERT INTO t_user_role (user_id, role_id)VALUES($1, $2)
		RETURNING id
		`, e.UserID, e.RoleID); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}
