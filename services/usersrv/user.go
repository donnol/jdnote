package usersrv

import (
	"github.com/donnol/jdnote/models/rolemodel"
	"github.com/donnol/jdnote/models/usermodel"
	"github.com/donnol/jdnote/models/userrolemodel"
	"github.com/donnol/jdnote/utils/context"
)

type userImpl struct {
	UserModel     usermodel.IUser
	UserRoleModel userrolemodel.IUserRole
}

// Check 检查
func (u *userImpl) Check(ctx context.Context) error {

	return nil
}

// GetByID 获取
func (u *userImpl) GetByID(ctx context.Context, id int) (e usermodel.Entity, err error) {
	return u.UserModel.GetByID(ctx, id)
}

// GetByName 获取
func (u *userImpl) GetByName(ctx context.Context, name string) (e usermodel.Entity, err error) {
	return u.UserModel.GetByName(ctx, name)
}

// GetFirst 获取首个用户
func (u *userImpl) GetFirst(ctx context.Context) (e usermodel.Entity, err error) {
	return u.UserModel.GetFirst(ctx)
}

// VerifyByNameAndPassword 校验用户密码
func (u *userImpl) VerifyByNameAndPassword(ctx context.Context, name, password string) (e usermodel.Entity, err error) {
	return u.UserModel.VerifyByNameAndPassword(ctx, name, password)
}

// Add 添加
func (u *userImpl) Add(ctx context.Context, e usermodel.Entity) (id int, err error) {

	// 用户模块添加
	if id, err = u.UserModel.Add(ctx, e); err != nil {
		return
	}

	// 用户角色模块添加
	ure := userrolemodel.Entity{
		UserID: id,
		RoleID: rolemodel.DefaultRoleID,
	}
	if _, err = u.UserRoleModel.Add(ctx, ure); err != nil {
		return
	}

	return
}