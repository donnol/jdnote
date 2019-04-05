package action

import (
	"testing"

	"github.com/donnol/jdnote/model"
	pg "github.com/donnol/jdnote/store/db/postgresql"
)

func TestGet(t *testing.T) {
	a := &Action{
		Base: model.Base{
			DB: pg.New(),
		},
		Action: "ALL",
	}
	if err := a.Add(); err != nil {
		t.Fatal(err)
	}

	if err := a.GetByID(8); err != nil {
		t.Fatal(err)
	}
	t.Log(a)
}
