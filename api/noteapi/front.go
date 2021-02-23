package noteapi

import (
	"context"

	"github.com/donnol/jdnote/services/notesrv"
	"github.com/donnol/jdnote/utils/route"
)

type Front struct {
	noteSrv notesrv.INote

	// 频率限制
	Limiter route.Limiter `rate:"Rate(0.25, 10)"`
}

func (front *Front) GetPage(ctx context.Context, p route.Param) (res route.Result, err error) {
	// 参数
	param := notesrv.PageParam{}
	if err = p.Parse(ctx, &param); err != nil {
		return
	}

	// 权限

	// 业务
	param.OnlyPublish = true
	result, err := front.noteSrv.GetPage(ctx, param)
	if err != nil {
		return
	}

	// 返回
	res.Data = result

	return
}

func (front *Front) Get(ctx context.Context, p route.Param) (res route.Result, err error) {
	// 参数
	param := notesrv.GetParam{}
	if err = p.Parse(ctx, &param); err != nil {
		return
	}

	// 业务
	mres, err := front.noteSrv.GetPublish(ctx, param.NoteID)
	if err != nil {
		return
	}

	// 返回
	res.Data = mres

	return
}
