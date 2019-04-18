package action

import (
	"testing"

	"github.com/donnol/jdnote/model"
)

func TestGet(t *testing.T) {
	a := &Action{
		Base: model.Base{
			DB: (&model.Base{}).New(),
		},
	}
	e := Entity{
		Action: "ALL",
	}
	var err error
	var id int
	if id, err = a.Add(e); err != nil {
		t.Fatal(err)
	} else if id == 0 {
		t.Fatal("Bad id")
	}

	if e, err := a.GetByID(id); err != nil {
		t.Fatal(err)
	} else if e.ID == 0 {
		t.Fatal("Bad entit")
	} else {
		t.Log(e)
	}
}
