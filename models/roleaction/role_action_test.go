package roleaction

import (
	"testing"

	"github.com/donnol/jdnote/app"
)

func TestCheckPerm(t *testing.T) {
	ra := &roleActionImpl{}
	ctx := app.DefaultCtx()

	if err := ra.CheckPerm(ctx, []string{"ALL"}); err != nil {
		t.Fatal(err)
	}
}
