package userrole

import (
	"testing"

	pg "github.com/donnol/jdnote/store/db/postgresql"
)

func TestGetByUserID(t *testing.T) {
	ur := &UserRole{
		DB: pg.New(),
	}
	if r, err := ur.GetByUserID(11); err != nil {
		t.Fatal(err)
	} else {
		t.Log(r)
	}
}
