package userrole

import (
	"log"
	"os"
	"testing"

	"github.com/donnol/jdnote/context"
	pg "github.com/donnol/jdnote/store/db/postgresql"
	utillog "github.com/donnol/jdnote/utils/log"
)

func TestGetByUserID(t *testing.T) {
	ur := &UserRole{}
	ctx := context.New((&pg.Base{}).New(), utillog.New(os.Stdout, "", log.LstdFlags), 0)
	if r, err := ur.GetByUserID(ctx, 38); err != nil {
		t.Fatal(err)
	} else {
		t.Log(r)
	}
}
