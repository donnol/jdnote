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

// Add 添加
func (n *Note) Add(ctx context.Context, p route.Param) (res route.Result, err error) {
	// 参数
	param := note.Param{}
	if err = p.Parse(&param); err != nil {
		return
	}

	// 权限
	if err = n.CheckLogin(ctx); err != nil {
		return
	}

	// 业务
	id, err := n.NoteAo.Add(ctx, param)
	if err != nil {
		return
	}

	// 返回
	res.Data = id

	return
}
