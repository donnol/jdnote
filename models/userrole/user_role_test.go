package userrole

import (
	"testing"

	"github.com/donnol/jdnote/app"
)

func TestGetByUserID(t *testing.T) {
	ur := &userRoleImpl{}
	ctx := app.DefaultCtx()
	if r, err := ur.GetByUserID(ctx, 38); err != nil {
		t.Fatal(err)
	} else {
		t.Log(r)
	}
}
