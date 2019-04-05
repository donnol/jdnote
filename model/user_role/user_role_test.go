package userrole

import (
	"testing"

	"github.com/donnol/jdnote/model"
	pg "github.com/donnol/jdnote/store/db/postgresql"
)

func TestGetByUserID(t *testing.T) {
	ur := &UserRole{
		Base: model.Base{
			DB: pg.New(),
		},
	}
	if r, err := ur.GetByUserID(38); err != nil {
		t.Fatal(err)
	} else {
		t.Log(r)
	}
}
