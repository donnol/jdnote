package auth

import (
	"net/http"

	"github.com/donnol/jdnote/route"
	"github.com/donnol/jdnote/service/user"
)

func init() {
	route.DefaultRouter.Register(http.MethodPost, "/login", &user.User{}, login)
	route.DefaultRouter.Register(http.MethodPost, "/add", &user.User{}, add)
}

func login(param route.Param) (r route.Result, err error) {
	// 参数
	p := param.RequestParam.(*user.User)

	// 权限
	_ = param.UserID

	// 业务
	if err := p.VerifyByNameAndPassword(p.Name, p.Password); err != nil {
		return r, err
	}
	r.CookieAfterLogin = p.ID

	return
}

func add(param route.Param) (r route.Result, err error) {
	// 参数
	p := param.RequestParam.(*user.User)

	// 权限
	_ = param.UserID

	// 业务
	if err := p.Add(); err != nil {
		return r, err
	}
	r.Data = p

	return
}
