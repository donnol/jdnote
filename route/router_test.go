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
}
