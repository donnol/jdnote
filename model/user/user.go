package user

import (
	"github.com/donnol/jdnote/model"
	"golang.org/x/crypto/bcrypt"
)

// User 用户
type User struct {
	model.Base
}

// GetByName 以名字获取用户
func (u *User) GetByName(name string) (e Entity, err error) {
	if err = u.DB.Get(&e, `SELECT id, name FROM t_user WHERE name = $1`, name); err != nil {
		return
	}

	return
}

// VerifyByNameAndPassword 以名字和密码校验用户
func (u *User) VerifyByNameAndPassword(name, password string) (e Entity, err error) {
	if err = u.DB.Get(&e, `SELECT id, name, password FROM t_user WHERE name = $1`, name); err != nil {
		return
	}
	u.Debugf("e: %+v\n", e)

	// 校验用户和密码
	if err = bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(password)); err != nil {
		return
	}

	return
}

// Add 添加
func (u *User) Add(e Entity) (id int, err error) {
	hashedPassword, err := u.hashPassword(e.Password)
	if err != nil {
		return
	}

	if err = u.DB.Get(&id, `INSERT INTO t_user (name, phone, email, password)
	VALUES($1, $2, $3, $4) RETURNING id`,
		e.Name,
		e.Phone,
		e.Email,
		hashedPassword,
	); err != nil {
		return
	}

	return
}
