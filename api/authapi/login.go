package authapi

import (
	"context"

	"github.com/donnol/jdnote/services/usersrv"
	utilctx "github.com/donnol/tools/context"
	"github.com/donnol/tools/route"
)

// Auth 认证
type Auth struct {
	// 方法配置-多个方法用逗号分隔
	Tx route.Method `tx:"AddUser"`

	// 使用的model
	UserSrv usersrv.IUser
}

// LoginUser 登录用户
type LoginUser struct {
	Name   string `json:"name"`
	Role   string `json:"role"`
	UserID int    `json:"userID"`
}

// GetIslogin 是否登录
func (auth *Auth) GetIslogin(ctx context.Context, param route.Param) (r route.Result, err error) {
	// 参数

	// 权限
	userID, err := utilctx.GetUserValue(ctx)
	if err != nil {
		return
	}
	if userID == 0 {
		r.Data = LoginUser{}
		return
	}

	// 业务
	u, err := auth.UserSrv.GetByID(ctx, userID)
	if err != nil {
		return
	}
	r.Data = LoginUser{
		Name:   u.Name,
		Role:   "admin",
		UserID: u.ID,
	}

	return
}

// AddLogin 登录
func (auth *Auth) AddLogin(ctx context.Context, param route.Param) (r route.Result, err error) {
	// 参数
	p := usersrv.Entity{}
	if err = param.Parse(ctx, &p); err != nil {
		return
	}

	// 权限

	// 业务
	re, err := auth.UserSrv.VerifyByNameAndPassword(ctx, p.Name, p.Password)
	if err != nil {
		return
	}
	r.CookieAfterLogin = re.ID
	r.Data = re

	return
}

// AddUser 添加用户
func (auth *Auth) AddUser(ctx context.Context, param route.Param) (r route.Result, err error) {
	// 参数
	p := usersrv.Entity{}
	if err = param.Parse(ctx, &p); err != nil {
		return
	}

	// 权限

	// 业务
	id, err := auth.UserSrv.Add(ctx, p)
	if err != nil {
		return
	}
	r.Data = id

	return
}

// GetUser 获取用户
func (auth *Auth) GetUser(ctx context.Context, param route.Param) (r route.Result, err error) {
	// 参数
	p := usersrv.Entity{}
	if err = param.Parse(ctx, &p); err != nil {
		return
	}

	// 权限

	// 业务
	re, err := auth.UserSrv.GetByName(ctx, p.Name)
	if err != nil {
		return
	}
	r.Data = re

	return
}
