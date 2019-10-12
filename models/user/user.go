package user

import (
	"github.com/donnol/jdnote/models"
	"github.com/donnol/jdnote/models/role/roledb"
	"github.com/donnol/jdnote/models/user/userdb"
	"github.com/donnol/jdnote/models/userrole/userroledb"
	"github.com/donnol/jdnote/utils/context"
)

// User 用户
type User struct {
	models.Base

	UserModel     userdb.User
	UserRoleModel userroledb.UserRole
}

// Check 检查
func (u *User) Check(ctx context.Context) error {

	return nil
}

// GetByID 获取
func (u *User) GetByID(ctx context.Context, id int) (e userdb.Entity, err error) {
	return u.UserModel.GetByID(ctx, id)
}

// GetByName 获取
func (u *User) GetByName(ctx context.Context, name string) (e userdb.Entity, err error) {
	return u.UserModel.GetByName(ctx, name)
}

// VerifyByNameAndPassword 校验用户密码
func (u *User) VerifyByNameAndPassword(ctx context.Context, name, password string) (e userdb.Entity, err error) {
	return u.UserModel.VerifyByNameAndPassword(ctx, name, password)
}

// Add 添加
func (u *User) Add(ctx context.Context, e userdb.Entity) (id int, err error) {

	// 用户模块添加
	if id, err = u.UserModel.Add(ctx, e); err != nil {
		return
	}

	// 用户角色模块添加
	ure := userroledb.Entity{
		UserID: id,
		RoleID: roledb.DefaultRoleID,
	}
	if _, err = u.UserRoleModel.Add(ctx, ure); err != nil {
		return
	}

	return
}
