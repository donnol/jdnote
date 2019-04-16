package auth

import (
	"github.com/donnol/jdnote/api"
	"github.com/donnol/jdnote/model/user"
	"github.com/donnol/jdnote/route"
	userao "github.com/donnol/jdnote/service/user"
)

func init() {
	route.DefaultRouter.Register(&Auth{})
}

// Auth 认证
type Auth struct {
	api.Base

	// 所属的Group
	V1 route.Group // 属于v1分组

	// 使用的model
	UserAo userao.User
}

// AddLogin 登录
func (auth *Auth) AddLogin(param route.Param) (r route.Result, err error) {
	// 参数
	p := user.Entity{}
	if err = param.Parse(&p); err != nil {
		return
	}

	// 权限
	_ = param.UserID

	// 业务
	var re user.Entity
	if re, err = auth.UserAo.VerifyByNameAndPassword(p.Name, p.Password); err != nil {
		return r, err
	}
	r.CookieAfterLogin = re.ID
	r.Data = re

	return
}

// AddUser 添加用户
func (auth *Auth) AddUser(param route.Param) (r route.Result, err error) {
	// 参数
	p := user.Entity{}
	if err = param.Parse(&p); err != nil {
		return
	}

	// 权限

	// 业务
	var id int
	if id, err = auth.UserAo.Add(p); err != nil {
		return r, err
	}
	r.Data = id

	return
}

// GetUser 获取用户
func (auth *Auth) GetUser(param route.Param) (r route.Result, err error) {
	// 参数
	p := user.Entity{}
	if err = param.Parse(&p); err != nil {
		return
	}

	// 权限
	_ = param.UserID

	// 业务
	var re user.Entity
	if re, err = auth.UserAo.GetByName(p.Name); err != nil {
		return
	}
	r.Data = re

	return
}
