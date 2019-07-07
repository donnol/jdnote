package action

import (
	"log"
	"os"
	"testing"

	"github.com/donnol/jdnote/context"
	pg "github.com/donnol/jdnote/store/db/postgresql"
	utillog "github.com/donnol/jdnote/utils/log"
)

func TestGet(t *testing.T) {
	a := &Action{}
	e := Entity{
		Action: "ALL",
	}
	var err error
	var id int
	ctx := context.New((&pg.Base{}).New(), utillog.New(os.Stdout, "", log.LstdFlags), 0)
	if id, err = a.Add(ctx, e); err != nil {
		t.Fatal(err)
	} else if id == 0 {
		t.Fatal("Bad id")
	}

	if e, err := a.GetByID(ctx, id); err != nil {
		t.Fatal(err)
	} else if e.ID == 0 {
		t.Fatal("Bad entit")
	} else {
		t.Log(e)
	}
}
