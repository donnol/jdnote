package route

import (
	"testing"

	pg "github.com/donnol/jdnote/store/db/postgresql"
)

// User 用户
type User struct {
	Model
	UserModel UserModel
}

// Model 模型
type Model struct {
	pg.DB
}

var (
	_ pg.DB = User{}
	_ pg.DB = Model{}
	_ pg.DB = UserModel{}
)

// UserModel 用户模块
type UserModel struct {
	Model
}

// GetNameByID 根据ID获取名字
func (um *UserModel) GetNameByID(id int) (string, error) {
	var name string
	if err := um.Get(&name, `SELECT name FROM t_user where id = $1`, id); err != nil {
		return name, err
	}
	return name, nil
}

func TestInitParamWithDB(t *testing.T) {
	r := initParamWithDB(&User{}, pg.New())
	t.Log(r)
	u := r.(*User)
	var id = 38
	var name string
	if err := u.Get(&name, `SELECT name FROM t_user where id = $1`, id); err != nil {
		t.Fatal(err)
	}
	t.Log(name)

	name, err := u.UserModel.GetNameByID(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(name)
}
