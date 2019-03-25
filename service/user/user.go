package userao

import (
	"fmt"

	"github.com/donnol/jdnote/model/role"
	"github.com/donnol/jdnote/model/user"
	userrole "github.com/donnol/jdnote/model/user_role"
	pg "github.com/donnol/jdnote/store/db/postgresql"
)

// User 用户
type User struct {
	user.User

	// TODO: 将需要用到的model添加进来
	UserModel     user.User
	UserRoleModel userrole.UserRole
}

// New 新建
func (u *User) New() interface{} {
	um := user.User{}
	um.DB = pg.New()

	return &User{
		User: um,
	}
}

// Check 检查
func (u *User) Check() error {
	if u.Password == "" {
		return fmt.Errorf("Empty Password")
	}

	return nil
}

// Add 添加
func (u *User) Add() error {
	if err := pg.WithTx(func(tx pg.DB) error {
		// 添加用户
		um := u.User
		um.DB = tx // TODO: 这里能不能不用显示赋值呢？
		if err := um.Add(); err != nil {
			return err
		}
		u.ID = um.ID

		// 添加角色
		ur := &userrole.UserRole{
			UserID: um.ID,
			RoleID: role.DefaultRoleID,
		}
		ur.DB = tx // 这里能不能不用显示赋值呢？
		if err := ur.Add(); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
