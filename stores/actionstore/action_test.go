package actionstore

import (
	"context"
	"testing"

	"github.com/donnol/jdnote/app"
	"github.com/donnol/jdnote/models/actionmodel"
)

func TestGet(t *testing.T) {
	a := &actionImpl{}
	e := actionmodel.Entity{
		Action: "ALL",
	}
	var err error
	var id int
	sctx := context.Background()
	_, ctx := app.New(sctx)
	if id, err = a.Add(ctx, e); err != nil {
		t.Fatal(err)
	} else if id == 0 {
		t.Fatal("Bad id")
	}

	if e, err := a.GetByID(ctx, id); err != nil {
		t.Fatal(err)
	} else if e.ID == 0 {
		t.Fatal("Bad entity")
	} else {
		t.Log(e)
	}
}
