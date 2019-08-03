package user

import (
	"log"
	"os"
	"testing"

	"github.com/donnol/jdnote/context"
	pg "github.com/donnol/jdnote/store/db/postgresql"
	utillog "github.com/donnol/jdnote/utils/log"
)

func TestGetByName(t *testing.T) {
	u := &User{}
	ctx := context.New((&pg.Base{}).New(), utillog.New(os.Stdout, "", log.LstdFlags), 0)
	if e, err := u.GetByName(ctx, "jd"); err != nil {
		t.Fatal(err)
	} else if e.ID == 0 {
		t.Fatal("Bad id")
	} else {
		t.Log(e)
	}
}

func TestAdd(t *testing.T) {
	u := &User{}
	e := Entity{
		Name:     "jd",
		Phone:    "jd",
		Email:    "jdlau@126.com",
		Password: "jd",
	}
	ctx := context.New((&pg.Base{}).New(), utillog.New(os.Stdout, "", log.LstdFlags), 0)
	if id, err := u.Add(ctx, e); err != nil {
		t.Fatal(err)
	} else if id == 0 {
		t.Fatal("Bad id")
	} else {
		t.Log(id)
	}
}
