package rolemodel

import (
	"github.com/donnol/jdnote/utils/context"
	"github.com/pkg/errors"
)

// DefaultRoleID 默认角色ID
var DefaultRoleID = 1

type roleImpl struct {
}

// GetByID 获取
func (r *roleImpl) GetByID(ctx context.Context, id int) (e Entity, err error) {
	if err = ctx.DB().GetContext(ctx, &e, `
		SELECT * FROM t_role WHERE id = $1
		`, id); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// Add 添加
func (r *roleImpl) Add(ctx context.Context, e Entity) (id int, err error) {
	if err = ctx.DB().GetContext(ctx, &id, `
		INSERT INTO t_role (role)VALUES($1) RETURNING id
		`, e.Role); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}
