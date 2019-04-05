package userrole

import (
	"github.com/donnol/jdnote/model"
)

// UserRole 用户角色
type UserRole struct {
	model.Base

	ID     int `json:"id"`                  // 记录ID
	UserID int `json:"userID" db:"user_id"` // 用户ID
	RoleID int `json:"roleID" db:"role_id"` // 角色ID
}

// GetByUserID 获取用户相关的角色
func (ur *UserRole) GetByUserID(userID int) ([]UserRole, error) {
	var result = make([]UserRole, 0)
	if err := ur.Select(&result, `
		SELECT * FROM t_user_role WHERE user_id = $1
		`, userID); err != nil {
		return result, err
	}

	return result, nil
}

// Add 添加
func (ur *UserRole) Add() error {
	var id int
	if err := ur.Get(&id, `
		INSERT INTO t_user_role (user_id, role_id)VALUES($1, $2)
		RETURNING id
		`, ur.UserID, ur.RoleID); err != nil {
		return err
	}
	ur.ID = id

	return nil
}
