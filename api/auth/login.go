package auth

import (
	"net/http"

	"github.com/donnol/jdnote/service/user"
	"github.com/donnol/jdnote/route"
)

func init() {
	route.DefaultRouter.Register(http.MethodGet, "/login", &user.User{}, login)
	route.DefaultRouter.Register(http.MethodPost, "/add", &user.User{}, add)
}

func login(param interface{}) (interface{}, error) {
	p := param.(*user.User)
	u := user.New()
	r := u.GetByName(p.Name)
	return r, nil
}

func add(param interface{}) (interface{}, error) {
	p := param.(*user.User)
	u := user.New()
	r := u.GetByName(p.Name)
	return r, nil
}