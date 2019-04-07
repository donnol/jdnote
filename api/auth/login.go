package auth

import (
	"github.com/donnol/jdnote/route"
	userao "github.com/donnol/jdnote/service/user"
)

func init() {
	// 用add/get/mod/del分别对应post/get/put/delete方法，路由从方法名(驼峰转换，如：getUser->get /user; getUserCurrent->get /user/current;)获取，但是参数一般都是每个接口都不同的，怎么设置好呢？
	// 还有，如login方法这种用post的，写成addLogin好像也不太好
	route.DefaultRouter.Register(&userao.User{}, addLogin)
	route.DefaultRouter.Register(&userao.User{}, addUser)
}

func addLogin(param route.Param) (r route.Result, err error) {
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

func addUser(param route.Param) (r route.Result, err error) {
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
