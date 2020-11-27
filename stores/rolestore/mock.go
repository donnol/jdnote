package rolestore

import (
	"github.com/donnol/jdnote/models/rolemodel"
	"github.com/donnol/jdnote/utils/context"
)

type RoleMock struct {
	AddFunc func(ctx context.Context, e rolemodel.Entity) (id int, err error)

	GetByIDFunc func(ctx context.Context, id int) (e rolemodel.Entity, err error)
}

var _ IRole = &RoleMock{}

func (mockRecv *RoleMock) Add(ctx context.Context, e rolemodel.Entity) (id int, err error) {
	return mockRecv.AddFunc(ctx, e)
}

func (mockRecv *RoleMock) GetByID(ctx context.Context, id int) (e rolemodel.Entity, err error) {
	return mockRecv.GetByIDFunc(ctx, id)
}
