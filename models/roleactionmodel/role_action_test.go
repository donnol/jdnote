package roleactionmodel

import (
	"context"
	"testing"

	"github.com/donnol/jdnote/app"
)

func TestCheckPerm(t *testing.T) {
	ra := &roleActionImpl{}
	sctx := context.Background()
	_, ctx := app.New(sctx)

	if err := ra.CheckPerm(ctx, []string{"ALL"}); err != nil {
		t.Fatal(err)
	}
}
