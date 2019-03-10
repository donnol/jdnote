package user

import (
	"fmt"

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

// Check 检查
func (u *User) Check() error {
	if u.Password == "" {
		return fmt.Errorf("Empty Password")
	}

	return nil
}
