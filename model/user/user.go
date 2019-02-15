package user

import (
	"github.com/donnol/jdnote/store/db"
	"github.com/jmoiron/sqlx"
)

// User 用户
type User struct {
	db.DB
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

// AddTwice 添加两次
func (u *User) AddTwice(name string) []int64 {
	var r = make([]int64, 0)

	if err := u.DB.WithTx(func(tx *sqlx.Tx) error {
		for i := 0; i < 2; i++ {
			var id int64
			err := tx.Get(&id, `INSERT INTO t_user (name, password) VALUES($1, '123') RETURNING id`, name)
			if err != nil {
				return err
			}
			r = append(r, id)
		}
		return nil
	}); err != nil {
		panic(err)
	}

	return r
}
