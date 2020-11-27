package notestore

import (
	"github.com/donnol/jdnote/models/notemodel"
	"github.com/donnol/jdnote/utils/common"
	"github.com/donnol/jdnote/utils/context"
)

type NoterMock struct {
	AddFunc func(ctx context.Context, entity notemodel.Entity) (id int, err error)

	AddOneFunc func(ctx context.Context) (id int, err error)

	DelFunc func(ctx context.Context, id int) (err error)

	GetFunc func(ctx context.Context, id int) (entity notemodel.Entity, err error)

	GetListFunc func(ctx context.Context, ids []int64) (entitys notemodel.EntityList, err error)

	GetPageFunc func(ctx context.Context, entity notemodel.Entity, param common.Param) (res notemodel.EntityList, total int, err error)

	ModFunc func(ctx context.Context, id int, entity notemodel.Entity) (err error)

	ModStatusFunc func(ctx context.Context, id int, status notemodel.Status) (err error)
}

var _ Noter = &NoterMock{}

func (mockRecv *NoterMock) Add(ctx context.Context, entity notemodel.Entity) (id int, err error) {
	return mockRecv.AddFunc(ctx, entity)
}

func (mockRecv *NoterMock) AddOne(ctx context.Context) (id int, err error) {
	return mockRecv.AddOneFunc(ctx)
}

func (mockRecv *NoterMock) Del(ctx context.Context, id int) (err error) {
	return mockRecv.DelFunc(ctx, id)
}

func (mockRecv *NoterMock) Get(ctx context.Context, id int) (entity notemodel.Entity, err error) {
	return mockRecv.GetFunc(ctx, id)
}

func (mockRecv *NoterMock) GetList(ctx context.Context, ids []int64) (entitys notemodel.EntityList, err error) {
	return mockRecv.GetListFunc(ctx, ids)
}

func (mockRecv *NoterMock) GetPage(ctx context.Context, entity notemodel.Entity, param common.Param) (res notemodel.EntityList, total int, err error) {
	return mockRecv.GetPageFunc(ctx, entity, param)
}

func (mockRecv *NoterMock) Mod(ctx context.Context, id int, entity notemodel.Entity) (err error) {
	return mockRecv.ModFunc(ctx, id, entity)
}

func (mockRecv *NoterMock) ModStatus(ctx context.Context, id int, status notemodel.Status) (err error) {
	return mockRecv.ModStatusFunc(ctx, id, status)
}
