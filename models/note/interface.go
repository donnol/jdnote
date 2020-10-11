package note

import (
	"github.com/donnol/jdnote/app"
	"github.com/donnol/jdnote/utils/common"
	"github.com/donnol/jdnote/utils/context"
)

// Noter 笔记接口
type Noter interface {
	AddOne(ctx context.Context) (id int, err error)
	Add(ctx context.Context, entity Entity) (id int, err error)
	Mod(ctx context.Context, id int, entity Entity) (err error)
	Del(ctx context.Context, id int) (err error)
	GetPage(ctx context.Context, entity Entity, param common.Param) (
		res EntityList,
		total int,
		err error,
	)
	Get(ctx context.Context, id int) (entity Entity, err error)
	GetList(ctx context.Context, ids []int64) (entitys EntityList, err error)
}

// New 新建
func New() Noter {
	return &noteImpl{}
}

func init() {
	app.MustRegisterProvider(New)
}
