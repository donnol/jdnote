package auth

import (
	"github.com/donnol/jdnote/api"
	"github.com/donnol/jdnote/models/user"
	"github.com/donnol/jdnote/models/user/userdb"
	"github.com/donnol/jdnote/route"
	"github.com/donnol/jdnote/utils/context"
)

func init() {
	route.Register(&Auth{})
}

// Auth 认证
type Auth struct {
	api.Base

	// 方法配置-多个方法用逗号分隔
	Tx route.Method `tx:"AddUser"`

	// 使用的model
	UserAo user.User
}

// LoginUser 登录用户
type LoginUser struct {
	Name   string `json:"name"`
	Role   string `json:"role"`
	UserID int    `json:"userID"`
}

// GetIslogin 是否登录
func (auth *Auth) GetIslogin(ctx context.Context, param route.Param) (r route.Result) {
	// 参数

	// 权限
	if ctx.UserID() == 0 {
		r.Data = LoginUser{}
		return
	}

	// 业务
	u, err := auth.UserAo.GetByID(ctx, ctx.UserID())
	r.SetErr(err)
	r.Data = LoginUser{
		Name:   u.Name,
		Role:   "admin",
		UserID: u.ID,
	}

	return
}

// AddLogin 登录
func (auth *Auth) AddLogin(ctx context.Context, param route.Param) (r route.Result) {
	// 参数
	p := userdb.Entity{}
	r.SetErr(param.Parse(&p))

	// 权限
	_ = ctx.UserID()

	// 业务
	re, err := auth.UserAo.VerifyByNameAndPassword(ctx, p.Name, p.Password)
	r.SetErr(err)
	r.CookieAfterLogin = re.ID
	r.Data = re

	return
}

// AddUser 添加用户
func (auth *Auth) AddUser(ctx context.Context, param route.Param) (r route.Result) {
	// 参数
	p := userdb.Entity{}
	r.SetErr(param.Parse(&p))

	// 权限

	// 业务
	id, err := auth.UserAo.Add(ctx, p)
	r.SetErr(err)
	r.Data = id

	return
}

// GetUser 获取用户
func (auth *Auth) GetUser(ctx context.Context, param route.Param) (r route.Result) {
	// 参数
	p := userdb.Entity{}
	r.SetErr(param.Parse(&p))

	// 权限
	_ = ctx.UserID()

	// 业务
	re, err := auth.UserAo.GetByName(ctx, p.Name)
	r.SetErr(err)
	r.Data = re

	return
}
