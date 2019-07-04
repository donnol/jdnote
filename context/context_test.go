package context

import (
	"log"
	"os"
	"testing"

	pg "github.com/donnol/jdnote/store/db/postgresql"
	utillog "github.com/donnol/jdnote/utils/log"
)

func TestContext(t *testing.T) {
	db := (&pg.Base{}).New()
	logger := utillog.New(os.Stdout, "[Context]", log.LstdFlags|log.Lshortfile)
	ctx := New(db, logger)
	var id uint
	if err := ctx.DB().Get(&id, "Select id from t_user order by id desc"); err != nil {
		t.Fatal(err)
	}
	ctx.Logger().Debugf("id: %+v\n", id)
}
