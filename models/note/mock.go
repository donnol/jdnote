package note

import (
	"github.com/donnol/jdnote/app"
	"github.com/donnol/jdnote/utils/context"
)

func init() {
	// 注入mock

	// mock case
}

type Mock struct {
	AddOneHandler  func(ctx context.Context) (id int, err error)
	AddHandler     func(ctx context.Context, entity Entity) (id int, err error)
	ModHandler     func(ctx context.Context, id int, entity Entity) (err error)
	DelHandler     func(ctx context.Context, id int) (err error)
	GetPageHandler func(ctx context.Context, entity Entity, param app.CommonParam) (
		res EntityList,
		total int,
		err error,
	)
	GetHandler     func(ctx context.Context, id int) (entity Entity, err error)
	GetListHandler func(ctx context.Context, ids []int64) (entitys EntityList, err error)
}

func (m Mock) AddOne(ctx context.Context) (id int, err error) {
	return m.AddOneHandler(ctx)
}

func (m Mock) Add(ctx context.Context, entity Entity) (id int, err error) {
	return m.AddHandler(ctx, entity)
}

func (m Mock) Mod(ctx context.Context, id int, entity Entity) (err error) {
	return m.ModHandler(ctx, id, entity)
}

func (m Mock) Del(ctx context.Context, id int) (err error) {
	return m.DelHandler(ctx, id)
}

func (m Mock) GetPage(ctx context.Context, entity Entity, param app.CommonParam) (
	res EntityList,
	total int,
	err error,
) {
	return m.GetPageHandler(ctx, entity, param)
}

func (m Mock) Get(ctx context.Context, id int) (entity Entity, err error) {
	return m.GetHandler(ctx, id)
}

func (m Mock) GetList(ctx context.Context, ids []int64) (entitys EntityList, err error) {
	return m.GetListHandler(ctx, ids)
}
