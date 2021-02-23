package actionstore

import (
	"context"
	"testing"

	"github.com/donnol/jdnote/models/actionmodel"
	"github.com/donnol/jdnote/utils/store/db"
)

func TestGet(t *testing.T) {
	var err error
	var id int

	ctx := context.Background()

	a := New(&db.DBMock{}) // 没有连接到数据库的
	e := actionmodel.Entity{
		Action: "WRITE",
	}

	if id, err = a.Add(ctx, e); err != nil {
		t.Fatal(err)
	} else if id == 0 {
		t.Fatal("Bad id")
	} else {
		t.Logf("id: %v\n", id)
	}

	if e, err := a.GetByID(ctx, id); err != nil {
		t.Fatal(err)
	} else if e.ID == 0 {
		t.Fatal("Bad entity")
	} else {
		t.Log(e)
	}
}
