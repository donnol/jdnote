package notestore

import (
	"context"
	"testing"

	"github.com/donnol/jdnote/models/notemodel"
	"github.com/donnol/jdnote/utils/common"
	"github.com/donnol/jdnote/utils/store/db"
)

func TestAddNote(t *testing.T) {
	ctx := context.Background()

	note := New(&db.DBMock{})

	// 加
	id, err := note.Add(ctx, notemodel.Entity{
		UserID: 1,
		Title:  "test",
		Detail: "test detail",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)

	// 查
	detail, err := note.Get(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(detail)

	// 分页
	r, err := note.GetPage(ctx, notemodel.Entity{}, common.DefaultParam)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)

	// 修改
	if err := note.Mod(ctx, id, &notemodel.Entity{
		Detail: "testDetail",
	}); err != nil {
		t.Fatal(err)
	}

	// 查
	detail, err = note.Get(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(detail)

	// 列表
	details, err := note.GetList(ctx, []int64{int64(id - 1), int64(id)})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(details)
}

func TestGetPage(t *testing.T) {
	ctx := context.Background()

	note := New(&db.DBMock{})

	// 分页
	r, err := note.GetPage(ctx, notemodel.Entity{}, common.Param{PageSize: 5})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestAddOne(t *testing.T) {
	ctx := context.Background()

	note := New(&db.DBMock{})

	id, err := note.AddOne(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}
