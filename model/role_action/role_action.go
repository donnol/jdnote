package roleaction

import (
	pg "github.com/donnol/jdnote/store/db/postgresql"
)

// RoleAction 角色动作
type RoleAction struct {
	pg.Base

	ID       int `json:"id"`
	RoleID   int `json:"roleID" db:"role_id"`
	ActionID int `json:"actionID" db:"action_id"`
}

// Add 添加
func (ra *RoleAction) Add() error {
	var id int
	if err := ra.DB.Get(&id, `
		INSERT INTO t_role_action (role_id, action_id)VALUES($1, $2)
		RETURNING id
		`, ra.RoleID, ra.ActionID); err != nil {
		return err
	}
	ra.ID = id

	return nil
}
