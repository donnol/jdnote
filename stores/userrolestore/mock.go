package userrolestore

import (
	"github.com/donnol/jdnote/models/userrolemodel"
	"github.com/donnol/jdnote/utils/context"
)

type UserRoleMock struct {
	AddFunc func(ctx context.Context, e userrolemodel.Entity) (id int, err error)

	GetByUserIDFunc func(ctx context.Context, userID int) (list []userrolemodel.Entity, err error)
}

var _ IUserRole = &UserRoleMock{}

func (mockRecv *UserRoleMock) Add(ctx context.Context, e userrolemodel.Entity) (id int, err error) {
	return mockRecv.AddFunc(ctx, e)
}

func (mockRecv *UserRoleMock) GetByUserID(ctx context.Context, userID int) (list []userrolemodel.Entity, err error) {
	return mockRecv.GetByUserIDFunc(ctx, userID)
}
