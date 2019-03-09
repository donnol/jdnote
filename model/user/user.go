package user

import (
	pg "github.com/donnol/jdnote/store/db/postgresql"
	"golang.org/x/crypto/bcrypt"
)

// User 用户
type User struct {
	ID       int    `json:"id" form:"id"`             // 记录ID
	Name     string `json:"name" form:"name"`         // 用户名
	Phone    string `json:"phone" form:"phone"`       // 手机号码
	Email    string `json:"email" form:"email"`       // 邮箱
	Password string `json:"password" form:"password"` // 密码
}

// GetByName 以名字获取用户
func (u *User) GetByName(name string) error {
	if err := pg.New().Get(u, `SELECT id, name FROM t_user WHERE name = $1`, name); err != nil {
		return err
	}

	return nil
}

// VerifyByNameAndPassword 以名字和密码校验用户
func (u *User) VerifyByNameAndPassword(name, password string) error {
	if err := pg.New().Get(u, `SELECT id, name, password FROM t_user WHERE name = $1`, name); err != nil {
		return err
	}

	// 校验用户和密码
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}

// Add 添加
func (u *User) Add() error {
	var id int

	hashedPassword, err := u.hashPassword(u.Password)
	if err != nil {
		return err
	}

	if err := pg.New().Get(&id, `INSERT INTO t_user (name, phone, email, password)
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

func (u *User) hashPassword(password string) ([]byte, error) {
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return []byte{}, err
	}

	return hashedPassword, nil
}
