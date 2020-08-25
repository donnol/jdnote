package auth

import (
	"github.com/donnol/jdnote/route"
	"github.com/donnol/jdnote/services/user"
	"github.com/donnol/jdnote/utils/context"
)

func init() {
	route.Register(&Auth{})
}

// Auth 认证
type Auth struct {
	// 方法配置-多个方法用逗号分隔
	Tx route.Method `tx:"AddUser"`

	// 使用的model
	UserAo user.IUser
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
	userID, err := context.GetUserValue(ctx)
	if err != nil {
		return
	}
	if userID == 0 {
		r.Data = LoginUser{}
		return
	}

	// 业务
	u, err := auth.UserAo.GetByID(ctx, userID)
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
	p := user.Entity{}
	if err = param.Parse(ctx, &p); err != nil {
		return
	}

	// 权限

	// 业务
	re, err := auth.UserAo.VerifyByNameAndPassword(ctx, p.Name, p.Password)
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
	p := user.Entity{}
	if err = param.Parse(ctx, &p); err != nil {
		return
	}

	// 权限

	// 业务
	id, err := auth.UserAo.Add(ctx, p)
	if err != nil {
		return
	}
	r.Data = id

	return
}

// GetUser 获取用户
func (auth *Auth) GetUser(ctx context.Context, param route.Param) (r route.Result, err error) {
	// 参数
	p := user.Entity{}
	if err = param.Parse(ctx, &p); err != nil {
		return
	}

	// 权限

	// 业务
	re, err := auth.UserAo.GetByName(ctx, p.Name)
	if err != nil {
		return
	}
	r.Data = re

	return
}
