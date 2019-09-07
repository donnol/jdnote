package roleaction

import (
	"testing"

	"github.com/donnol/jdnote/utils/context"
	pg "github.com/donnol/jdnote/store/db/postgresql"
)

func TestCheckPerm(t *testing.T) {
	ra := &RoleAction{}
	ctx := context.New((&pg.Base{}).New(), nil, 114)

	if err := ra.CheckPerm(ctx, []string{"ALL"}); err != nil {
		t.Fatal(err)
	}
}
