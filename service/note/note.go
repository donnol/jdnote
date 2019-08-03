package note

import (
	"strconv"

	"github.com/donnol/jdnote/context"
	"github.com/donnol/jdnote/model"
	"github.com/donnol/jdnote/model/note"
)

// Note 笔记
type Note struct {
	NoteModel note.Note
}

// GetPage 获取分页
func (n *Note) GetPage(ctx context.Context, entity PageParam, param model.CommonParam) (r PageResult, err error) {
	res, total, err := n.NoteModel.GetPage(ctx, note.Entity{}, param)
	if err != nil {
		return
	}
	r.Total = total

	var tmp Result
	for _, single := range res {
		tmp = Result{}

		tmp.UserName = strconv.Itoa(single.UserID)
		tmp.Title = single.Title
		tmp.Detail = single.Detail
		tmp.CreatedAt = single.CreatedAt.Unix()

		r.List = append(r.List, tmp)
	}

	return
}

// AddOne 添加
func (n *Note) AddOne(ctx context.Context) (id int, err error) {
	id, err = n.NoteModel.AddOne(ctx)
	if err != nil {
		return
	}

	return
}

// Mod 修改
func (n *Note) Mod(ctx context.Context, id int, p Param) (err error) {
	if err = n.NoteModel.Mod(ctx, id, note.Entity{
		Title:  p.Title,
		Detail: p.Detail,
	}); err != nil {
		return
	}

	return
}
