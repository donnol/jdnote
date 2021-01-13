package userrolestore

import (
	"context"
	"testing"

	"github.com/donnol/jdnote/internal/initializers"
)

func TestGetByUserID(t *testing.T) {
	sctx := context.Background()
	_, ctx := initializers.New(sctx)

	ur := New()
	if r, err := ur.GetByUserID(ctx, 1); err != nil {
		t.Fatal(err)
	} else {
		t.Log(r)
	}
}
