package role

import (
	pg "github.com/donnol/jdnote/store/db/postgresql"
)

// DefaultRoleID 默认角色ID
var DefaultRoleID = 1

// Role 角色
type Role struct {
	pg.Base

	ID   int    `json:"id"`   // 记录ID
	Role string `json:"role"` // 角色
}

// Get 获取
func (r *Role) Get(id int) error {
	if err := r.DB.Get(r, `
		SELECT * FROM t_role WHERE id = $1
		`, id); err != nil {
		return err
	}

	return nil
}

// Add 添加
func (r *Role) Add() error {
	var id int
	if err := r.DB.Get(&id, `
		INSERT INTO t_role (role)VALUES($1) RETURNING id
		`, r.Role); err != nil {
		return err
	}
	r.ID = id

	return nil
}
