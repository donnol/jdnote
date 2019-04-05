package role

import (
	"github.com/donnol/jdnote/model"
)

// DefaultRoleID 默认角色ID
var DefaultRoleID = 1

// Role 角色
type Role struct {
	model.Base

	ID   int    `json:"id"`   // 记录ID
	Role string `json:"role"` // 角色
}

// GetByID 获取
func (r *Role) GetByID(id int) error {
	if err := r.Get(r, `
		SELECT * FROM t_role WHERE id = $1
		`, id); err != nil {
		return err
	}

	return nil
}

// Add 添加
func (r *Role) Add() error {
	var id int
	if err := r.Get(&id, `
		INSERT INTO t_role (role)VALUES($1) RETURNING id
		`, r.Role); err != nil {
		return err
	}
	r.ID = id

	return nil
}
