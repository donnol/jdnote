package roleaction

import (
	"fmt"

	"github.com/donnol/jdnote/context"
	"github.com/donnol/jdnote/model"
	"github.com/lib/pq"
)

// RoleAction 角色动作
type RoleAction struct {
	model.Base
}

// Add 添加
func (ra *RoleAction) Add(ctx context.Context, e Entity) (id int, err error) {
	if err = ctx.DB().GetContext(ctx, &id, `
		INSERT INTO t_role_action (role_id, action_id)VALUES($1, $2)
		RETURNING id
		`, e.RoleID, e.ActionID); err != nil {
		return
	}

	return
}

// CheckPerm 检查权限
func (ra *RoleAction) CheckPerm(ctx context.Context, perms []string) error {
	var exist bool
	if err := ctx.DB().GetContext(ctx, &exist, `
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
		ctx.UserID(),
		pq.StringArray(perms),
	); err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("No permission")
	}

	return nil
}
