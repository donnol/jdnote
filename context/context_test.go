package context

import (
	"log"
	"os"
	"testing"

	pg "github.com/donnol/jdnote/store/db/postgresql"
	utillog "github.com/donnol/jdnote/utils/log"
	"github.com/jmoiron/sqlx"
)

func TestContext(t *testing.T) {
	db := (&pg.Base{}).New()
	logger := utillog.New(os.Stdout, "[Context]", log.LstdFlags|log.Lshortfile)
	userID := 1
	ctx := New(db, logger, userID)
	var id uint
	if err := ctx.DB().Get(&id, "Select id from t_user order by id desc"); err != nil {
		t.Fatal(err)
	}
	ctx.Logger().Debugf("id: %+v\n", id)
}

func TestContextDB(t *testing.T) {
	ctx := Default()

	// get 1
	ctxDB := ctx.DB()
	ctx.Logger().Debugf("%p\n", ctxDB)

	// set
	db := ctxDB.(*sqlx.DB)
	*db = sqlx.DB{}

	// get 2
	ctxDB1 := ctx.DB()
	ctx.Logger().Debugf("%p\n", ctxDB1)
}
