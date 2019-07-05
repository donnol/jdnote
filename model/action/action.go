package action

import (
	"github.com/donnol/jdnote/context"
	"github.com/donnol/jdnote/model"
)

// Action 操作
type Action struct {
	model.Base
}

// GetByID 获取
func (a *Action) GetByID(ctx context.Context, id int) (e Entity, err error) {
	if err = ctx.DB().Get(&e, `
		SELECT * FROM t_action WHERE id = $1
		`, id); err != nil {
		return
	}

	return
}

// Add 添加
func (a *Action) Add(ctx context.Context, e Entity) (id int, err error) {
	if err = ctx.DB().Get(&id, `
		INSERT INTO t_action (action)VALUES($1) RETURNING id
		`, e.Action); err != nil {
		return
	}

	return
}
