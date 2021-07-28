package noteapi

import (
	"context"

	"github.com/donnol/jdnote/services/authsrv"
	"github.com/donnol/jdnote/services/notesrv"
	"github.com/donnol/tools/route"
)

// Note 笔记
type Note struct {
	authSrv authsrv.IAuth

	noteSrv notesrv.INote

	// 频率限制
	Limiter route.Limiter `rate:"Rate(0.25, 2)"`
}

// GetPage 获取分页
func (n *Note) GetPage(ctx context.Context, p route.Param) (res route.Result, err error) {
	// 参数
	param := notesrv.PageParam{}
	if err = p.Parse(ctx, &param); err != nil {
		return
	}

	// 权限
	if err = n.authSrv.CheckLogin(ctx); err != nil {
		return
	}

	// 业务
	result, err := n.noteSrv.GetPage(ctx, param)
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
	param := notesrv.GetParam{}
	if err = p.Parse(ctx, &param); err != nil {
		return
	}

	// 权限
	if err = n.authSrv.CheckLogin(ctx); err != nil {
		return
	}

	// 业务
	mres, err := n.noteSrv.Get(ctx, param.NoteID)
	if err != nil {
		return
	}

	// 返回
	res.Data = mres

	return
}

// Add 添加
func (n *Note) Add(ctx context.Context, p route.Param) (res route.Result, err error) {
	// 参数

	// 权限
	if err = n.authSrv.CheckLogin(ctx); err != nil {
		return
	}

	// 业务
	id, err := n.noteSrv.AddOne(ctx)
	if err != nil {
		return
	}

	// 返回
	res.Data = route.AddResult{
		ID: id,
	}

	return
}

// Mod 修改
func (n *Note) Mod(ctx context.Context, p route.Param) (res route.Result, err error) {
	// 参数
	param := notesrv.ModParam{}
	if err = p.Parse(ctx, &param); err != nil {
		return
	}

	// 权限
	if err = n.authSrv.CheckLogin(ctx); err != nil {
		return
	}

	// 业务
	err = n.noteSrv.Mod(ctx, param.NoteID, &param.Param)
	if err != nil {
		return
	}

	// 返回

	return
}

func (n *Note) ModPublish(ctx context.Context, p route.Param) (res route.Result, err error) {
	// 参数
	param := notesrv.GetParam{}
	if err = p.Parse(ctx, &param); err != nil {
		return
	}

	// 权限
	if err = n.authSrv.CheckLogin(ctx); err != nil {
		return
	}

	// 业务
	err = n.noteSrv.Publish(ctx, param.NoteID)
	if err != nil {
		return
	}

	// 返回

	return
}

func (n *Note) ModHide(ctx context.Context, p route.Param) (res route.Result, err error) {
	// 参数
	param := notesrv.GetParam{}
	if err = p.Parse(ctx, &param); err != nil {
		return
	}

	// 权限
	if err = n.authSrv.CheckLogin(ctx); err != nil {
		return
	}

	// 业务
	err = n.noteSrv.Hide(ctx, param.NoteID)
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
	if err = p.Parse(ctx, &param); err != nil {
		return
	}

	// 权限
	if err = n.authSrv.CheckLogin(ctx); err != nil {
		return
	}

	// 业务
	err = n.noteSrv.Del(ctx, param.NoteID)
	if err != nil {
		return
	}

	// 返回

	return
}
