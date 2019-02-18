package user

import "testing"

func TestGetByName(t *testing.T) {
	u := &User{}
	user := u.GetByName("jd")
	t.Log(user)
}

func TestAdd(t *testing.T) {
	u := &User{
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
