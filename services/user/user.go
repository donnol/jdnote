package user

import (
	"github.com/donnol/jdnote/models"
	"github.com/donnol/jdnote/models/role"
	"github.com/donnol/jdnote/models/user"
	"github.com/donnol/jdnote/models/userrole"
	"github.com/donnol/jdnote/utils/context"
)

// User 用户
type User struct {
	models.Base

	UserModel     user.User
	UserRoleModel userrole.UserRole
}

// Check 检查
func (u *User) Check(ctx context.Context) error {

	return nil
}

// GetByID 获取
func (u *User) GetByID(ctx context.Context, id int) (e user.Entity, err error) {
	return u.UserModel.GetByID(ctx, id)
}

// GetByName 获取
func (u *User) GetByName(ctx context.Context, name string) (e user.Entity, err error) {
	return u.UserModel.GetByName(ctx, name)
}

// GetFirst 获取首个用户
func (u *User) GetFirst(ctx context.Context) (e user.Entity, err error) {
	return u.UserModel.GetFirst(ctx)
}

// VerifyByNameAndPassword 校验用户密码
func (u *User) VerifyByNameAndPassword(ctx context.Context, name, password string) (e user.Entity, err error) {
	return u.UserModel.VerifyByNameAndPassword(ctx, name, password)
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
