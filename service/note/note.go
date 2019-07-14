package note

import (
	"github.com/donnol/jdnote/context"
	"github.com/donnol/jdnote/model/note"
)

// Note 笔记
type Note struct {
	NoteModel note.Note
}

// Add 添加
func (n *Note) Add(ctx context.Context, p Param) (id int, err error) {
	id, err = n.NoteModel.Add(ctx, note.Entity{
		UserID: ctx.UserID(),
		Title:  p.Title,
		Detail: p.Detail,
	})
	if err != nil {
		return
	}

	return
}
