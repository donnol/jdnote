package roleactionstore

import (
	"context"
	"testing"

	utilctx "github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/store/db"
)

func TestCheckPerm(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, utilctx.UserKey, 1)

	ra := New(&db.DBMock{})
	if err := ra.CheckPerm(ctx, []string{"ALL"}); err != nil {
		t.Fatal(err)
	}
}
