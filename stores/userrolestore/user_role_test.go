package userrolestore

import (
	"context"
	"testing"

	"github.com/donnol/jdnote/internal/initializers"
)

func TestGetByUserID(t *testing.T) {
	ur := &userRoleImpl{}
	sctx := context.Background()
	_, ctx := initializers.New(sctx)
	if r, err := ur.GetByUserID(ctx, 38); err != nil {
		t.Fatal(err)
	} else {
		t.Log(r)
	}
}
