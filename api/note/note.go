package note

import (
	"github.com/donnol/jdnote/api"
	"github.com/donnol/jdnote/context"
	"github.com/donnol/jdnote/route"
	"github.com/donnol/jdnote/service/note"
)

func init() {
	route.Register(&Note{})
}

// Note 笔记
type Note struct {
	api.Base

	NoteAo note.Note
}

// GetPage 获取分页
func (n *Note) GetPage(ctx context.Context, p route.Param) (res route.Result, err error) {
	// 参数
	param := note.PageParam{}
	if err = p.Parse(&param); err != nil {
		return
	}

	// 权限
	if err = n.CheckLogin(ctx); err != nil {
		return
	}

	// 业务
	result, err := n.NoteAo.GetPage(ctx, param)
	if err != nil {
		return
	}

	// 返回
	res.Data = result

	return
}

// Get 获取
func (n *Note) Get(ctx context.Context, p route.Param) (res route.Result, err error) {
	// 参数
	param := struct {
		NoteID int `json:"noteID"`
	}{}
	if err = p.Parse(&param); err != nil {
		return
	}

	// 权限
	if err = n.CheckLogin(ctx); err != nil {
		return
	}

	// 业务
	mres := n.NoteAo.Get2(ctx, param.NoteID)
	result, err := mres.Data(), mres.Err()
	if err != nil {
		return
	}

	// 返回
	res.Data = result

	return
}

// Add 添加
func (n *Note) Add(ctx context.Context, p route.Param) (res route.Result, err error) {
	// 参数

	// 权限
	if err = n.CheckLogin(ctx); err != nil {
		return
	}

	// 业务
	id, err := n.NoteAo.AddOne(ctx)
	if err != nil {
		return
	}

	// 返回
	res.Data = api.AddResult{
		ID: id,
	}

	return
}

// Mod 修改
func (n *Note) Mod(ctx context.Context, p route.Param) (res route.Result, err error) {
	// 参数
	param := note.ModParam{}
	if err = p.Parse(&param); err != nil {
		return
	}

	// 权限
	if err = n.CheckLogin(ctx); err != nil {
		return
	}

	// 业务
	err = n.NoteAo.Mod(ctx, param.NoteID, param.Param)
	if err != nil {
		return
	}

	// 返回

	return
}

// Del 删除
func (n *Note) Del(ctx context.Context, p route.Param) (res route.Result, err error) {
	// 参数
	param := struct {
		NoteID int `json:"noteID"`
	}{}
	if err = p.Parse(&param); err != nil {
		return
	}

	// 权限
	if err = n.CheckLogin(ctx); err != nil {
		return
	}

	// 业务
	err = n.NoteAo.Del(ctx, param.NoteID)
	if err != nil {
		return
	}

	// 返回

	return
}
