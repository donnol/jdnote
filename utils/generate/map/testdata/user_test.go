package user

import "testing"

func TestUserSliceToMap(t *testing.T) {
	for _, cas := range []struct {
		UserList []User
	}{
		{[]User{User{Name: "jd"}}},
	} {
		m := UserSliceToMap(cas.UserList)
		t.Log(m)
	}
}
