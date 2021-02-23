package roleactionstore

import (
	"context"

	"github.com/donnol/jdnote/models/roleactionmodel"
)

type RoleActionMock struct {
	AddFunc func(ctx context.Context, e roleactionmodel.Entity) (id int, err error)

	CheckPermFunc func(ctx context.Context, perms []string) error
}

var _ IRoleAction = &RoleActionMock{}

func (mockRecv *RoleActionMock) Add(ctx context.Context, e roleactionmodel.Entity) (id int, err error) {
	return mockRecv.AddFunc(ctx, e)
}

func (mockRecv *RoleActionMock) CheckPerm(ctx context.Context, perms []string) error {
	return mockRecv.CheckPermFunc(ctx, perms)
}
