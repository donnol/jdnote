package userao

import (
	"github.com/donnol/jdnote/context"
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
func (u *User) Add(ctx context.Context, e user.Entity) (id int, err error) {

	// 如果ctx.DB()返回的不是一个事务，那么怎么在这里开启一个事务呢？
	// 下面的写法也可以，但是如果里面的方法也开启了事务呢？怎么确保这里的是同一个事务呢？
	// if err = (&pg.Base{}).WithTx(func(tx pg.DB) error {
	// 	txCtx := context.New(tx, ctx.Logger(), ctx.UserID())

	// 	u.UserModel.Add(txCtx, e)
	// 	u.UserRoleModel.Add(txCtx, userrole.Entity{})

	// 	return nil
	// }); err != nil {
	// 	return
	// }

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
