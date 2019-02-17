package user

import "testing"

func TestUser(t *testing.T) {
	for _, cas := range []struct {
		In []User
	}{
		{[]User{{1, "jd"}, {2, "lau"}}},
	} {
		r := QueryColumn(cas.In)
		t.Log(r)
	}
}
