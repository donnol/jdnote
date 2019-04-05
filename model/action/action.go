package action

import (
	"github.com/donnol/jdnote/model"
)

// Action 操作
type Action struct {
	model.Base

	ID     int    `json:"id"`     // 记录ID
	Action string `json:"action"` // 操作
}

// GetByID 获取
func (a *Action) GetByID(id int) error {
	if err := a.Get(a, `
		SELECT * FROM t_action WHERE id = $1
		`, id); err != nil {
		return err
	}

	return nil
}

// Add 添加
func (a *Action) Add() error {
	var id int
	if err := a.Get(&id, `
		INSERT INTO t_action (action)VALUES($1) RETURNING id
		`, a.Action); err != nil {
		return err
	}
	a.ID = id

	return nil
}
