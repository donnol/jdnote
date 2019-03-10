package user

import (
	"testing"

	pg "github.com/donnol/jdnote/store/db/postgresql"
)

func TestGetByName(t *testing.T) {
	u := &User{
		DB: pg.New(),
	}
	if err := u.GetByName("jd"); err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestAdd(t *testing.T) {
	u := &User{
		DB:       pg.New(),
		Name:     "jd",
		Phone:    "13420693396",
		Email:    "jdlau@126.com",
		Password: "13420693396",
	}
	if err := u.Add(); err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}
