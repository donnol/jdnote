package note

import (
	"github.com/donnol/jdnote/models/note"
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/inject"
)

type INote interface {
	GetPage(ctx context.Context, param PageParam) (r PageResult, err error)
	Get(ctx context.Context, id int) (r Result, err error)
	AddOne(ctx context.Context) (id int, err error)
	Mod(ctx context.Context, id int, p Param) (err error)
	Del(ctx context.Context, id int) (err error)
	Publish(ctx context.Context, id int) error
	Hide(ctx context.Context, id int) error
}

func New(
	nm note.Noter,
) INote {
	return &noteImpl{
		noteModel: nm,
	}
}

func init() {
	inject.MustRegisterProvider(New)
}
