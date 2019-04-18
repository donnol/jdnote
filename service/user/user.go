package userao

import (
	"github.com/donnol/jdnote/model/role"
	"github.com/donnol/jdnote/model/user"
	userrole "github.com/donnol/jdnote/model/user_role"
)

// User 用户
type User struct {
	user.User

	UserModel     user.User
	UserRoleModel userrole.UserRole
}

// Check 检查
func (u *User) Check() error {

	return nil
}

// Add 添加
func (u *User) Add(e user.Entity) (id int, err error) {
	if err = u.InjectTx(u, func(v interface{}) error {
		// 使用断言转换具体类型
		nu := v.(*User)

		// 用户模块添加
		if id, err = nu.UserModel.Add(e); err != nil {
			return err
		}

		// 用户角色模块添加
		ure := userrole.Entity{
			UserID: id,
			RoleID: role.DefaultRoleID,
		}
		if _, err = nu.UserRoleModel.Add(ure); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return
	}

	return
}
