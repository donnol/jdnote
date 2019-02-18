package user

import (
	pg "github.com/donnol/jdnote/store/db/postgresql"
	"golang.org/x/crypto/bcrypt"
)

// User 用户
type User struct {
	pg.DB

	ID       int    `json:"id" form:"id"`       // 记录ID
	Name     string `json:"name" form:"name"`   // 用户名
	Phone    string `json:"phone" form:"phone"` // 手机号码
	Email    string `json:"email" form:"email"` // 邮箱
	Password string `json:"-" form:"password"`  // 密码
}

// GetByName 以名字获取用户
func (u *User) GetByName(name string) User {
	var r User
	if err := u.DB.New().Get(&r, `SELECT id, name FROM t_user WHERE name = $1`, name); err != nil {
		panic(err)
	}
	return r
}

// Add 添加
func (u *User) Add() error {
	var id int

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(u.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	if err := u.DB.New().Get(&id, `INSERT INTO t_user (name, phone, email, password)
		VALUES($1, $2, $3, $4) RETURNING id`,
		u.Name,
		u.Phone,
		u.Email,
		hashedPassword,
	); err != nil {
		return err
	}
	u.ID = id

	return nil
}
