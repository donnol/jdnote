package userrole

import (
	"github.com/donnol/jdnote/context"
	"github.com/donnol/jdnote/model"
	"github.com/pkg/errors"
)

// UserRole 用户角色
type UserRole struct {
	model.Base
}

// GetByUserID 获取用户相关的角色
func (ur *UserRole) GetByUserID(ctx context.Context, userID int) (list []Entity, err error) {
	if err = ctx.DB().SelectContext(ctx, &list, `
		SELECT * FROM t_user_role WHERE user_id = $1
		`, userID); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// Add 添加
func (ur *UserRole) Add(ctx context.Context, e Entity) (id int, err error) {
	if err = ctx.DB().GetContext(ctx, &id, `
		INSERT INTO t_user_role (user_id, role_id)VALUES($1, $2)
		RETURNING id
		`, e.UserID, e.RoleID); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}
