package roleactiondata

import (
	"testing"

	"github.com/donnol/jdnote/models"
)

func TestCheckPerm(t *testing.T) {
	ra := &RoleAction{}
	ctx := models.DefaultCtx()

	if err := ra.CheckPerm(ctx, []string{"ALL"}); err != nil {
		t.Fatal(err)
	}
}
