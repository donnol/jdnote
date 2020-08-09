package user

import (
	"fmt"

	"github.com/donnol/jdnote/utils/context"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type userImpl struct {
}

// GetByID 以id获取用户
func (u *userImpl) GetByID(ctx context.Context, id int) (e Entity, err error) {
	if err = ctx.DB().GetContext(ctx, &e, `SELECT id, name FROM t_user WHERE id = $1`, id); err != nil {
		err = errors.WithMessage(err, fmt.Sprintf("id: %d", id))
		return
	}

	return
}

// GetByName 以名字获取用户
func (u *userImpl) GetByName(ctx context.Context, name string) (e Entity, err error) {
	if err = ctx.DB().GetContext(ctx, &e, `SELECT id, name FROM t_user WHERE name = $1`, name); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// GetFirst 获取首个用户
func (u *userImpl) GetFirst(ctx context.Context) (e Entity, err error) {
	if err = ctx.DB().GetContext(ctx, &e, `SELECT id, name FROM t_user ORDER BY id ASC LIMIT 1`); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// VerifyByNameAndPassword 以名字和密码校验用户
func (u *userImpl) VerifyByNameAndPassword(ctx context.Context, name, password string) (e Entity, err error) {
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
func (u *userImpl) Add(ctx context.Context, e Entity) (id int, err error) {
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
