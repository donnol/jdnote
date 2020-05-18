package note

import "github.com/donnol/jdnote/utils/context"

type Noter interface {
	GetPage(ctx context.Context, param PageParam) (r PageResult, err error)
	Get(ctx context.Context, id int) (r Result, err error)
	AddOne(ctx context.Context) (id int, err error)
	Mod(ctx context.Context, id int, p Param) (err error)
	Del(ctx context.Context, id int) (err error)
	Publish(ctx context.Context, id int) error
	Hide(ctx context.Context, id int) error
}
