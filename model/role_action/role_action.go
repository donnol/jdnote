package roleaction

import (
	"github.com/donnol/jdnote/model"
)

// RoleAction 角色动作
type RoleAction struct {
	model.Base
}

// Add 添加
func (ra *RoleAction) Add(e Entity) (id int, err error) {
	if err = ra.DB.Get(&id, `
		INSERT INTO t_role_action (role_id, action_id)VALUES($1, $2)
		RETURNING id
		`, e.RoleID, e.ActionID); err != nil {
		return
	}

	return
}
