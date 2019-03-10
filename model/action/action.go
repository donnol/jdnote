package action

import (
	pg "github.com/donnol/jdnote/store/db/postgresql"
)

// Action 操作
type Action struct {
	pg.DB `json:"-" db:"-"`

	ID     int    `json:"id"`     // 记录ID
	Action string `json:"action"` // 操作
}

// Get 获取
func (a *Action) Get(id int) error {
	if err := a.DB.Get(a, `
		SELECT * FROM t_action WHERE id = $1
		`, id); err != nil {
		return err
	}

	return nil
}

// Add 添加
func (a *Action) Add() error {
	var id int
	if err := a.DB.Get(&id, `
		INSERT INTO t_action (action)VALUES($1) RETURNING id
		`, a.Action); err != nil {
		return err
	}
	a.ID = id

	return nil
}
