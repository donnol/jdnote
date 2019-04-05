package roleaction

import (
	"github.com/donnol/jdnote/model"
)

// RoleAction 角色动作
type RoleAction struct {
	model.Base

	ID       int `json:"id"`                      // 记录ID
	RoleID   int `json:"roleID" db:"role_id"`     // 角色ID
	ActionID int `json:"actionID" db:"action_id"` // 动作ID
}

// Add 添加
func (ra *RoleAction) Add() error {
	var id int
	if err := ra.Get(&id, `
		INSERT INTO t_role_action (role_id, action_id)VALUES($1, $2)
		RETURNING id
		`, ra.RoleID, ra.ActionID); err != nil {
		return err
	}
	ra.ID = id

	return nil
}
