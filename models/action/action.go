package action

import (
	"github.com/donnol/jdnote/models"
	"github.com/donnol/jdnote/utils/context"
	"github.com/pkg/errors"
)

// Action 操作
type Action struct {
	models.Base
}

// GetByID 获取
func (a *Action) GetByID(ctx context.Context, id int) (e Entity, err error) {
	if err = ctx.DB().GetContext(ctx, &e, `
		SELECT * FROM t_action WHERE id = $1
		`, id); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// Add 添加
func (a *Action) Add(ctx context.Context, e Entity) (id int, err error) {
	if err = ctx.DB().GetContext(ctx, &id, `
		INSERT INTO t_action (action)VALUES($1) RETURNING id
		`, e.Action); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}
