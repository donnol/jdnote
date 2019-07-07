package note

import (
	"testing"

	"github.com/donnol/jdnote/context"
	"github.com/donnol/jdnote/model"
)

func TestAddNote(t *testing.T) {
	note := &Note{}
	ctx := context.Default()

	// 加
	id, err := note.AddNote(ctx, Entity{
		UserID: 1,
		Title:  "test",
		Detail: "test detail",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)

	// 查
	detail, err := note.GetNote(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(detail)

	// 列表查
	r, err := note.GetNoteList(ctx, Entity{}, model.DefaultCommonParam)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)

	// 修改
	if err := note.ModifyNote(ctx, id, Entity{
		Detail: "testDetail",
	}); err != nil {
		t.Fatal(err)
	}

	// 查
	detail, err = note.GetNote(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(detail)
}
