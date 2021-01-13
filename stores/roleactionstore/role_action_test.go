package roleactionstore

import (
	"context"
	"testing"

	"github.com/donnol/jdnote/internal/initializers"
)

func TestCheckPerm(t *testing.T) {
	ra := &roleActionImpl{}
	sctx := context.Background()
	_, ctx := initializers.New(sctx)

	if err := ra.CheckPerm(ctx, []string{"ALL"}); err != nil {
		t.Fatal(err)
	}
}
