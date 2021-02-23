package actionstore

import (
	"context"

	"github.com/donnol/jdnote/models/actionmodel"
)

type ActionMock struct {
	AddFunc func(ctx context.Context, e actionmodel.Entity) (id int, err error)

	GetByIDFunc func(ctx context.Context, id int) (e actionmodel.Entity, err error)
}

var _ IAction = &ActionMock{}

func (mockRecv *ActionMock) Add(ctx context.Context, e actionmodel.Entity) (id int, err error) {
	return mockRecv.AddFunc(ctx, e)
}

func (mockRecv *ActionMock) GetByID(ctx context.Context, id int) (e actionmodel.Entity, err error) {
	return mockRecv.GetByIDFunc(ctx, id)
}
