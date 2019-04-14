package action

import (
	"github.com/donnol/jdnote/model"
)

// Action 操作
type Action struct {
	model.Base
}

// GetByID 获取
func (a *Action) GetByID(id int) (e Entity, err error) {
	if err = a.DB.Get(&e, `
		SELECT * FROM t_action WHERE id = $1
		`, id); err != nil {
		return
	}

	return
}

// Add 添加
func (a *Action) Add(e Entity) (id int, err error) {
	if err = a.DB.Get(&id, `
		INSERT INTO t_action (action)VALUES($1) RETURNING id
		`, e.Action); err != nil {
		return
	}

	return
}
