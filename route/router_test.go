package route

import (
	"testing"

	pg "github.com/donnol/jdnote/store/db/postgresql"
)

// User 用户
type User struct {
	pg.DB
}

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
