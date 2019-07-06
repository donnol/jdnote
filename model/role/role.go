package role

import (
	"github.com/donnol/jdnote/context"
	"github.com/donnol/jdnote/model"
)

// DefaultRoleID 默认角色ID
var DefaultRoleID = 1

// Role 角色
type Role struct {
	model.Base
}

// GetByID 获取
func (r *Role) GetByID(ctx context.Context, id int) (e Entity, err error) {
	if err = ctx.DB().GetContext(ctx, &e, `
		SELECT * FROM t_role WHERE id = $1
		`, id); err != nil {
		return
	}

	return
}

// Add 添加
func (r *Role) Add(ctx context.Context, e Entity) (id int, err error) {
	if err = ctx.DB().GetContext(ctx, &id, `
		INSERT INTO t_role (role)VALUES($1) RETURNING id
		`, e.Role); err != nil {
		return
	}

	return
}
