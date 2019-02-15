package user

import "testing"

func TestGetByName(t *testing.T) {
	u := &User{}
	user := u.GetByName("jd")
	t.Log(user)
}

func TestAddTwice(t *testing.T) {
	u := &User{}
	r := u.AddTwice("jd")
	t.Log(r)
}
