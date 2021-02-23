package actionstore

import (
	"context"

	"github.com/donnol/jdnote/models/actionmodel"
	"github.com/donnol/jdnote/utils/store/db"
	"github.com/pkg/errors"
)

type actionImpl struct {
	db db.DB
}

// GetByID 获取
func (a *actionImpl) GetByID(ctx context.Context, id int) (e actionmodel.Entity, err error) {
	if err = db.DBFromCtxValue(ctx, a.db).GetContext(ctx, &e, `
		SELECT * FROM t_action WHERE id = $1
		`, id); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// Add 添加
func (a *actionImpl) Add(ctx context.Context, e actionmodel.Entity) (id int, err error) {
	if err = db.DBFromCtxValue(ctx, a.db).GetContext(ctx, &id, `
		INSERT INTO t_action (action)VALUES($1) RETURNING id
		`, e.Action); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}
