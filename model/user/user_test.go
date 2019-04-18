package user

import (
	"testing"

	"github.com/donnol/jdnote/model"
)

func TestGetByName(t *testing.T) {
	u := &User{
		Base: model.Base{
			DB: (&model.Base{}).New(),
		},
	}
	if e, err := u.GetByName("jd"); err != nil {
		t.Fatal(err)
	} else if e.ID == 0 {
		t.Fatal("Bad id")
	} else {
		t.Log(e)
	}
}

func TestAdd(t *testing.T) {
	u := &User{
		Base: model.Base{
			DB: (&model.Base{}).New(),
		},
	}
	e := Entity{
		Name:     "jd",
		Phone:    "13420693396",
		Email:    "jdlau@126.com",
		Password: "13420693396",
	}
	if id, err := u.Add(e); err != nil {
		t.Fatal(err)
	} else if id == 0 {
		t.Fatal("Bad id")
	} else {
		t.Log(id)
	}
}
