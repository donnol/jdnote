package userrole

import (
	"github.com/donnol/jdnote/model"
)

// UserRole 用户角色
type UserRole struct {
	model.Base
}

// GetByUserID 获取用户相关的角色
func (ur *UserRole) GetByUserID(userID int) (list []Entity, err error) {
	if err = ur.Select(&list, `
		SELECT * FROM t_user_role WHERE user_id = $1
		`, userID); err != nil {
		return
	}

	return
}

// Add 添加
func (ur *UserRole) Add(e Entity) (id int, err error) {
	if err = ur.Get(&id, `
		INSERT INTO t_user_role (user_id, role_id)VALUES($1, $2)
		RETURNING id
		`, e.UserID, e.RoleID); err != nil {
		return
	}

	return
}
