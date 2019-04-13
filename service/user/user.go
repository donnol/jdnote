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

	UserModel     user.User
	UserRoleModel userrole.UserRole
}

// Check 检查
func (u *User) Check() error {

	return nil
}

// Add 添加
func (u *User) Add(e user.Entity) (id int, err error) {
	if err = pg.WithTx(func(tx pg.DB) error {
		var err error

		// 添加用户-必须获取model的副本，这样才不会改变model的DB值
		um := u.UserModel
		um.SetTx(tx)
		// 如果像这样直接调用SetTx，就会改变model里的DB值，对后面的操作会一直有影响
		// u.UserModel.SetTx(tx)
		if id, err = um.Add(e); err != nil {
			return err
		}

		// 添加角色
		ur := u.UserRoleModel
		ur.SetTx(tx)
		ure := userrole.Entity{
			UserID: id,
			RoleID: role.DefaultRoleID,
		}
		if _, err = ur.Add(ure); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return
	}

	return
}
