package action

import (
	"testing"

	"github.com/donnol/jdnote/models"
)

func TestGet(t *testing.T) {
	a := &actionImpl{}
	e := Entity{
		Action: "ALL",
	}
	var err error
	var id int
	ctx := models.DefaultCtx()
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
