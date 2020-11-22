package notesrv

import "github.com/donnol/jdnote/utils/context"

type Mock struct {
	GetPageHandler    func(ctx context.Context, param PageParam) (r PageResult, err error)
	GetHandler        func(ctx context.Context, id int) (r Result, err error)
	GetPublishHandler func(ctx context.Context, id int) (r Result, err error)
	AddOneHandler     func(ctx context.Context) (id int, err error)
	ModHandler        func(ctx context.Context, id int, p Param) (err error)
	DelHandler        func(ctx context.Context, id int) (err error)
	PublishHandler    func(ctx context.Context, id int) error
	HideHandler       func(ctx context.Context, id int) error
}

var (
	_ INote = Mock{}
)

func (m Mock) GetPage(ctx context.Context, param PageParam) (r PageResult, err error) {
	return m.GetPageHandler(ctx, param)
}

func (m Mock) Get(ctx context.Context, id int) (r Result, err error) {
	return m.GetHandler(ctx, id)
}

func (m Mock) GetPublish(ctx context.Context, id int) (r Result, err error) {
	return m.GetPublishHandler(ctx, id)
}

func (m Mock) AddOne(ctx context.Context) (id int, err error) {
	return m.AddOneHandler(ctx)
}

func (m Mock) Mod(ctx context.Context, id int, p Param) (err error) {
	return m.ModHandler(ctx, id, p)
}

func (m Mock) Del(ctx context.Context, id int) (err error) {
	return m.DelHandler(ctx, id)
}

func (m Mock) Publish(ctx context.Context, id int) error {
	return m.PublishHandler(ctx, id)
}

func (m Mock) Hide(ctx context.Context, id int) error {
	return m.HideHandler(ctx, id)
}
