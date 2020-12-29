package notesrv

import (
	"github.com/donnol/jdnote/stores/notestore"
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/timer"
)

type INote interface {
	GetPage(ctx context.Context, param PageParam) (r PageResult, err error)
	Get(ctx context.Context, id int) (r Result, err error)
	GetPublish(ctx context.Context, id int) (r Result, err error)
	AddOne(ctx context.Context) (id int, err error)
	Mod(ctx context.Context, id int, p *Param) (err error)
	Del(ctx context.Context, id int) (err error)
	Publish(ctx context.Context, id int) error
	Hide(ctx context.Context, id int) error
	Timer(ctx context.Context) timer.FuncJob
}

func New(
	nm notestore.Noter,
) INote {
	return &noteImpl{
		noteStore: nm,
	}
}
