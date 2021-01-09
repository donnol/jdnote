package notestore

import (
	"github.com/donnol/jdnote/models/notemodel"
	"github.com/donnol/jdnote/utils/common"
	"github.com/donnol/jdnote/utils/context"
)

// Noter 笔记接口
type Noter interface {
	AddOne(ctx context.Context) (id int, err error)
	Add(ctx context.Context, entity notemodel.Entity) (id int, err error)
	Mod(ctx context.Context, id int, entity *notemodel.Entity) (err error)
	ModStatus(ctx context.Context, id int, status notemodel.Status) (err error)
	Del(ctx context.Context, id int) (err error)
	GetPage(ctx context.Context, entity notemodel.Entity, param common.Param) (
		res []notemodel.EntityWithTotal,
		err error,
	)
	Get(ctx context.Context, id int) (entity notemodel.Entity, err error)
	GetList(ctx context.Context, ids []int64) (entitys []notemodel.Entity, err error)
}

// New 新建
func New() Noter {
	return &noteImpl{}
}
