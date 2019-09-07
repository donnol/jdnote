package noteao

import (
	"strconv"

	"github.com/donnol/jdnote/models"
	"github.com/donnol/jdnote/models/noteao/notedb"
	"github.com/donnol/jdnote/utils/context"
)

// Note 笔记
type Note struct {
	NoteModel notedb.Note
}

// GetPage 获取分页
func (n *Note) GetPage(ctx context.Context, param PageParam) (r PageResult, err error) {
	entity := notedb.Entity{
		Title:  param.Title,
		Detail: param.Detail,
	}
	res, total, err := n.NoteModel.GetPage(ctx, entity, param.CommonParam)
	if err != nil {
		return
	}
	r.Total = total

	var tmp Result
	for _, single := range res {
		tmp = Result{}

		tmp.NoteID = single.ID
		tmp.UserName = strconv.Itoa(single.UserID)
		tmp.Title = single.Title
		tmp.Detail = single.Detail
		tmp.CreatedAt = single.CreatedAt.Unix()

		r.List = append(r.List, tmp)
	}

	return
}

// Get 获取
func (n *Note) Get(ctx context.Context, id int) (r Result, err error) {
	e, err := n.NoteModel.Get(ctx, id)
	if err != nil {
		return
	}
	r.NoteID = e.ID
	r.Title = e.Title
	r.Detail = e.Detail
	r.UserName = strconv.Itoa(e.UserID)
	r.CreatedAt = e.CreatedAt.Unix()

	return
}

// Get2 获取
func (n *Note) Get2(ctx context.Context, id int) (res models.Result) {
	var r Result

	// 从model拿数据
	e := n.NoteModel.Get2(ctx, id).Unwrap().(notedb.Entity)

	// 处理具体数据
	r.NoteID = e.ID
	r.Title = e.Title
	r.Detail = e.Detail
	r.UserName = strconv.Itoa(e.UserID)
	r.CreatedAt = e.CreatedAt.Unix()

	// 设置数据
	res.SetData(r)

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
	if err = n.NoteModel.Mod(ctx, id, notedb.Entity{
		Title:  p.Title,
		Detail: p.Detail,
	}); err != nil {
		return
	}

	return
}

// Del 删除
func (n *Note) Del(ctx context.Context, id int) (err error) {
	err = n.NoteModel.Del(ctx, id)
	if err != nil {
		return
	}

	return
}
