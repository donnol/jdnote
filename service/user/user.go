package userao

import (
	"github.com/donnol/jdnote/model/role"
	"github.com/donnol/jdnote/model/user"
	userrole "github.com/donnol/jdnote/model/user_role"
	pg "github.com/donnol/jdnote/store/db/postgresql"
)

// User 用户
type User struct {
	user.User

	UserModel     user.User         `json:"-"`
	UserRoleModel userrole.UserRole `json:"-"`
}

// Check 检查
func (u *User) Check() error {

	return nil
}

// Add 添加
func (u *User) Add() error {
	if err := pg.WithTx(func(tx pg.DB) error {
		// 添加用户-必须获取model的副本，这样才不会改变model的DB值
		um := u.UserModel
		um.SetTx(tx)
		// 如果像这样直接调用SetTx，就会改变model里的DB值，对后面的操作会一直有影响
		// u.UserModel.SetTx(tx)
		if err := um.Add(); err != nil {
			return err
		}
		u.ID = um.ID

		// 添加角色
		ur := u.UserRoleModel
		ur.UserID = um.ID
		ur.RoleID = role.DefaultRoleID
		ur.SetTx(tx)
		if err := ur.Add(); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
