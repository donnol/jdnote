package auth

import (
	"net/http"

	"github.com/donnol/jdnote/route"
	userao "github.com/donnol/jdnote/service/user"
)

func init() {
	route.DefaultRouter.Register(http.MethodPost, "/login", &userao.User{}, login)
	route.DefaultRouter.Register(http.MethodPost, "/add", &userao.User{}, add)
}

// TODO: login在这个文件的三个地方出现了，能不能将其简化为一个呢？
func login(param route.Param) (r route.Result, err error) {
	// 参数
	p := param.RequestParam.(*userao.User)

	// 权限
	_ = param.UserID

	// 业务
	if err := p.VerifyByNameAndPassword(p.Name, p.Password); err != nil {
		return r, err
	}
	r.CookieAfterLogin = p.ID
	p.Password = ""
	r.Data = p

	return
}

func add(param route.Param) (r route.Result, err error) {
	// 参数
	p := param.RequestParam.(*userao.User)

	// 权限
	_ = param.UserID

	// 业务
	if err := p.Add(); err != nil {
		return r, err
	}
	r.Data = p

	return
}
