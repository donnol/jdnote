package userroledata

import (
	"testing"

	"github.com/donnol/jdnote/models"
)

func TestGetByUserID(t *testing.T) {
	ur := &UserRole{}
	ctx := models.DefaultCtx()
	if r, err := ur.GetByUserID(ctx, 38); err != nil {
		t.Fatal(err)
	} else {
		t.Log(r)
	}
}
