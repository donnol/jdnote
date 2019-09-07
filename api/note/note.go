package note

import (
	"github.com/donnol/jdnote/api"
	"github.com/donnol/jdnote/utils/context"
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
func (n *Note) GetPage(ctx context.Context, p route.Param) (res route.Result) {
	// 参数
	param := note.PageParam{}
	res.SetErr(p.Parse(&param))

	// 权限
	res.SetErr(n.CheckLogin(ctx))

	// 业务
	result, err := n.NoteAo.GetPage(ctx, param)
	res.SetErr(err)

	// 返回
	res.Data = result

	return
}

// Get 获取
func (n *Note) Get(ctx context.Context, p route.Param) (res route.Result) {
	// 参数
	param := struct {
		NoteID int `json:"noteID"`
	}{}
	res.SetErr(p.Parse(&param))

	// 权限
	res.SetErr(n.CheckLogin(ctx))

	// 业务
	mres := n.NoteAo.Get2(ctx, param.NoteID)
	res.SetErr(mres.Err())

	// 返回
	res.Data = mres.Data()

	return
}

// Add 添加
func (n *Note) Add(ctx context.Context, p route.Param) (res route.Result) {
	// 参数

	// 权限
	res.SetErr(n.CheckLogin(ctx))

	// 业务
	id, err := n.NoteAo.AddOne(ctx)
	res.SetErr(err)

	// 返回
	res.Data = api.AddResult{
		ID: id,
	}

	return
}

// Mod 修改
func (n *Note) Mod(ctx context.Context, p route.Param) (res route.Result) {
	// 参数
	param := note.ModParam{}
	res.SetErr(p.Parse(&param))

	// 权限
	res.SetErr(n.CheckLogin(ctx))

	// 业务
	err := n.NoteAo.Mod(ctx, param.NoteID, param.Param)
	res.SetErr(err)

	// 返回

	return
}

// Del 删除
func (n *Note) Del(ctx context.Context, p route.Param) (res route.Result) {
	// 参数
	param := struct {
		NoteID int `json:"noteID"`
	}{}
	res.SetErr(p.Parse(&param))

	// 权限
	res.SetErr(n.CheckLogin(ctx))

	// 业务
	err := n.NoteAo.Del(ctx, param.NoteID)
	res.SetErr(err)

	// 返回

	return
}
