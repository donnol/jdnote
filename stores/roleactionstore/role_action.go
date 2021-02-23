package roleactionstore

import (
	"context"

	"github.com/donnol/jdnote/models/roleactionmodel"
	utilctx "github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/store/db"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type roleActionImpl struct {
	db db.DB
}

// Add 添加
func (ra *roleActionImpl) Add(ctx context.Context, e roleactionmodel.Entity) (id int, err error) {
	if err = db.DBFromCtxValue(ctx, ra.db).GetContext(ctx, &id, `
		INSERT INTO t_role_action (role_id, action_id)VALUES($1, $2)
		RETURNING id
		`, e.RoleID, e.ActionID); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// CheckPerm 检查权限
func (ra *roleActionImpl) CheckPerm(ctx context.Context, perms []string) error {
	var exist bool
	if err := db.DBFromCtxValue(ctx, ra.db).GetContext(ctx, &exist, `
		select exists(
			select * from 
			t_role_action ra
			left join t_role r on r.id = ra.role_id
			left join t_user_role ur on ur.role_id = r.id
			left join t_action a on a.id = ra.action_id
			where true
			and ur.user_id = $1
			and a.action = any($2)
		)
		`,
		utilctx.MustGetUserValue(ctx),
		pq.StringArray(perms),
	); err != nil {
		err = errors.WithStack(err)
		return err
	}
	if !exist {
		return errors.Errorf("No permission")
	}

	return nil
}
