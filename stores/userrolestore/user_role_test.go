package userrolestore

import (
	"context"
	"testing"

	"github.com/donnol/jdnote/utils/store/db"
)

func TestGetByUserID(t *testing.T) {
	ctx := context.Background()

	ur := New(&db.DBMock{})
	if r, err := ur.GetByUserID(ctx, 1); err != nil {
		t.Fatal(err)
	} else {
		t.Log(r)
	}
}
