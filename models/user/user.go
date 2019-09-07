package user

import (
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// User 用户
type User struct {
	model.Base
}

// GetByID 以id获取用户
func (u *User) GetByID(ctx context.Context, id int) (e Entity, err error) {
	if err = ctx.DB().GetContext(ctx, &e, `SELECT id, name FROM t_user WHERE id = $1`, id); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// GetByName 以名字获取用户
func (u *User) GetByName(ctx context.Context, name string) (e Entity, err error) {
	if err = ctx.DB().GetContext(ctx, &e, `SELECT id, name FROM t_user WHERE name = $1`, name); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// VerifyByNameAndPassword 以名字和密码校验用户
func (u *User) VerifyByNameAndPassword(ctx context.Context, name, password string) (e Entity, err error) {
	if err = ctx.DB().GetContext(ctx, &e, `SELECT id, name, password FROM t_user WHERE name = $1`, name); err != nil {
		err = errors.WithStack(err)
		return
	}

	// 校验用户和密码
	if err = bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(password)); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// Add 添加
func (u *User) Add(ctx context.Context, e Entity) (id int, err error) {
	hashedPassword, err := u.hashPassword(e.Password)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if err = ctx.DB().GetContext(ctx, &id, `INSERT INTO t_user (name, phone, email, password)
	VALUES($1, $2, $3, $4) RETURNING id`,
		e.Name,
		e.Phone,
		e.Email,
		hashedPassword,
	); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}
