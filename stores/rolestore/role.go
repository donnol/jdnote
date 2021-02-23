package rolestore

import (
	"context"

	"github.com/donnol/jdnote/models/rolemodel"
	"github.com/donnol/jdnote/utils/store/db"
	"github.com/pkg/errors"
)

type roleImpl struct {
	db db.DB
}

// GetByID 获取
func (r *roleImpl) GetByID(ctx context.Context, id int) (e rolemodel.Entity, err error) {
	if err = db.DBFromCtxValue(ctx, r.db).GetContext(ctx, &e, `
		SELECT * FROM t_role WHERE id = $1
		`, id); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// Add 添加
func (r *roleImpl) Add(ctx context.Context, e rolemodel.Entity) (id int, err error) {
	if err = db.DBFromCtxValue(ctx, r.db).GetContext(ctx, &id, `
		INSERT INTO t_role (role)VALUES($1) RETURNING id
		`, e.Role); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}
