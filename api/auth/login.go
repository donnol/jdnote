package auth

import (
	"github.com/donnol/jdnote/api"
	"github.com/donnol/jdnote/route"
	userao "github.com/donnol/jdnote/service/user"
)

func init() {
	// 用add/get/mod/del分别对应post/get/put/delete方法，路由从方法名(驼峰转换，如：getUser->get /user; getUserCurrent->get /user/current;)获取，但是参数一般都是每个接口都不同的，怎么设置好呢？
	// 还有，如login方法这种用post的，写成addLogin好像也不太好
	route.DefaultRouter.RegisterStruct(&Auth{})
}

// Auth 认证
type Auth struct {
	api.Base

	UserAo userao.User
}

// AddLogin 登录
func (auth *Auth) AddLogin() (r route.Result, err error) {
	// 参数 TODO: 怎么获取呢？
	p := &userao.User{}

	// 权限

	// 业务
	if err := auth.UserAo.VerifyByNameAndPassword(p.Name, p.Password); err != nil {
		return r, err
	}
	r.CookieAfterLogin = p.ID
	p.Password = ""
	r.Data = p

	return
}

// AddUser 添加用户
func (auth *Auth) AddUser() (r route.Result, err error) {
	// 参数
	p := &userao.User{}

	// 权限

	// 业务
	if err := auth.UserAo.Add(); err != nil {
		return r, err
	}
	r.Data = p

	return
}
