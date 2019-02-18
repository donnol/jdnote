package user

import (
	pg "github.com/donnol/jdnote/store/db/postgresql"
)

// User 用户
type User struct {
	pg.DB
	ID   int    `json:"id" form:"id"`     // 记录ID
	Name string `json:"name" form:"name"` // 用户名
}

// GetByName 以名字获取用户
func (u *User) GetByName(name string) User {
	var r User
	if err := u.DB.New().Get(&r, `SELECT id, name FROM t_user WHERE name = $1`, name); err != nil {
		panic(err)
	}
	return r
}
