package user

import (
	"testing"

	"github.com/donnol/jdnote/app"
)

func TestGetByName(t *testing.T) {
	u := &userImpl{}
	ctx := app.DefaultCtx()
	if e, err := u.GetByName(ctx, "jd"); err != nil {
		t.Fatal(err)
	} else if e.ID == 0 {
		t.Fatal("Bad id")
	} else {
		t.Log(e)
	}
}

func TestAdd(t *testing.T) {
	u := &userImpl{}
	e := Entity{
		Name:     "jd",
		Phone:    "jd",
		Email:    "jdlau@126.com",
		Password: "jd",
	}
	ctx := app.DefaultCtx()
	if id, err := u.Add(ctx, e); err != nil {
		t.Fatal(err)
	} else if id == 0 {
		t.Fatal("Bad id")
	} else {
		t.Log(id)
	}
}
