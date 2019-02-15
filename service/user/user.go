package user

import (
	"github.com/donnol/jdnote/model/user"
)

// User 用户
type User struct {
	user.User
}

// New 新建
func New() *User {
	return &User{
		User: user.User{},
	}
}
