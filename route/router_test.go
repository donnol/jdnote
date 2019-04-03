package route

import (
	"testing"

	pg "github.com/donnol/jdnote/store/db/postgresql"
)

// User 用户
type User struct {
	Model
}

// Model 模型
type Model struct {
	pg.DB
}

var (
	_ pg.DB = User{}
	_ pg.DB = Model{}
)

func TestInitParamWithDB(t *testing.T) {
	r := initParamWithDB(&User{}, pg.New())
	t.Log(r)
	u := r.(*User)
	var name string
	if err := u.Get(&name, `SELECT name FROM t_user where id = $1`, 38); err != nil {
		t.Fatal(err)
	}
	t.Log(name)
}
