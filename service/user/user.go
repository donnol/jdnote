package userao

import (
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/models/role"
	"github.com/donnol/jdnote/models/user"
	userrole "github.com/donnol/jdnote/models/user_role"
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
func (u *User) Add(ctx context.Context, e user.Entity) (id int, err error) {

	// 用户模块添加
	if id, err = u.UserModel.Add(ctx, e); err != nil {
		return
	}

	// 用户角色模块添加
	ure := userrole.Entity{
		UserID: id,
		RoleID: role.DefaultRoleID,
	}
	if _, err = u.UserRoleModel.Add(ctx, ure); err != nil {
		return
	}

	return
}
