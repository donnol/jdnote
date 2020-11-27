package userstore

import (
	"context"
	"testing"

	"github.com/donnol/jdnote/app"
	"github.com/donnol/jdnote/models/usermodel"
)

func TestGetByName(t *testing.T) {
	u := &userImpl{}
	sctx := context.Background()
	_, ctx := app.New(sctx)
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
	e := usermodel.Entity{
		Name:     "jd",
		Phone:    "jd",
		Email:    "jdlau@126.com",
		Password: "jd",
	}
	sctx := context.Background()
	_, ctx := app.New(sctx)
	if id, err := u.Add(ctx, e); err != nil {
		t.Fatal(err)
	} else if id == 0 {
		t.Fatal("Bad id")
	} else {
		t.Log(id)
	}
}
