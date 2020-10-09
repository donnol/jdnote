package note

import (
	"context"
	"testing"

	"github.com/donnol/jdnote/app"
)

func TestAddNote(t *testing.T) {
	note := &noteImpl{}
	sctx := context.Background()
	_, ctx := app.New(sctx)

	// 加
	id, err := note.Add(ctx, Entity{
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
	r, total, err := note.GetPage(ctx, Entity{}, app.DefaultCommonParam)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r, total)

	// 修改
	if err := note.Mod(ctx, id, Entity{
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
	note := &noteImpl{}
	sctx := context.Background()
	_, ctx := app.New(sctx)

	// 分页
	r, total, err := note.GetPage(ctx, Entity{}, app.CommonParam{PageSize: 5})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r, total)
}

func TestAddOne(t *testing.T) {
	note := &noteImpl{}
	sctx := context.Background()
	_, ctx := app.New(sctx)

	id, err := note.AddOne(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}
