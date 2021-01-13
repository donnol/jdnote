package roleactionstore

import (
	stdctx "context"
	"testing"

	"github.com/donnol/jdnote/internal/initializers"
	"github.com/donnol/jdnote/utils/context"
)

func TestCheckPerm(t *testing.T) {
	sctx := stdctx.Background()
	_, ctx := initializers.New(sctx)
	ctx = context.WithValue(ctx, context.UserKey, 1)

	ra := New()
	if err := ra.CheckPerm(ctx, []string{"ALL"}); err != nil {
		t.Fatal(err)
	}
}
