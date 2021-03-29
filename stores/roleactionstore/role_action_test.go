package roleactionstore

import (
	"context"
	"testing"

	"github.com/donnol/jdnote/utils/store/db"
	utilctx "github.com/donnol/tools/context"
)

func TestCheckPerm(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, utilctx.UserKey, 1)

	ra := New(&db.DBMock{})
	if err := ra.CheckPerm(ctx, []string{"ALL"}); err != nil {
		t.Fatal(err)
	}
}
