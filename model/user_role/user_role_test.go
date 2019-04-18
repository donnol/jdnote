package userrole

import (
	"testing"

	"github.com/donnol/jdnote/model"
)

func TestGetByUserID(t *testing.T) {
	ur := &UserRole{
		Base: model.Base{
			DB: (&model.Base{}).New(),
		},
	}
	if r, err := ur.GetByUserID(38); err != nil {
		t.Fatal(err)
	} else {
		t.Log(r)
	}
}
