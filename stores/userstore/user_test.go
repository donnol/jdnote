package userstore

import (
	"context"
	"testing"

	"github.com/donnol/jdnote/internal/initializers"
	"github.com/donnol/jdnote/models/usermodel"
)

func TestGetByName(t *testing.T) {
	sctx := context.Background()
	_, ctx := initializers.New(sctx)

	u := New()
	if e, err := u.GetByName(ctx, "jd"); err != nil {
		t.Fatal(err)
	} else if e.ID == 0 {
		t.Fatal("Bad id")
	} else {
		t.Log(e)
	}
}

func TestAdd(t *testing.T) {
	sctx := context.Background()
	_, ctx := initializers.New(sctx)

	u := New()
	e := usermodel.Entity{
		Name:     "jd",
		Phone:    "jd",
		Email:    "jdlau@126.com",
		Password: "jd",
	}
	if id, err := u.Add(ctx, e); err != nil {
		t.Fatal(err)
	} else if id == 0 {
		t.Fatal("Bad id")
	} else {
		t.Log(id)
	}
}
