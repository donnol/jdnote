package actionstore

import (
	"context"
	"testing"

	"github.com/donnol/jdnote/internal/initializers"
	"github.com/donnol/jdnote/models/actionmodel"
)

func TestGet(t *testing.T) {
	var err error
	var id int

	sctx := context.Background()
	_, ctx := initializers.New(sctx)

	a := New()
	e := actionmodel.Entity{
		Action: "WRITE",
	}

	if id, err = a.Add(ctx, e); err != nil {
		t.Fatal(err)
	} else if id == 0 {
		t.Fatal("Bad id")
	} else {
		t.Logf("id: %v\n", id)
	}

	if e, err := a.GetByID(ctx, id); err != nil {
		t.Fatal(err)
	} else if e.ID == 0 {
		t.Fatal("Bad entity")
	} else {
		t.Log(e)
	}
}
