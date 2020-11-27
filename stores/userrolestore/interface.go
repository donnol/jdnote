package userrolestore

import (
	"github.com/donnol/jdnote/models/userrolemodel"
	"github.com/donnol/jdnote/utils/context"
)

type IUserRole interface {
	Add(ctx context.Context, e userrolemodel.Entity) (id int, err error)
	GetByUserID(ctx context.Context, userID int) (list []userrolemodel.Entity, err error)
}

func New() IUserRole {
	return &userRoleImpl{}
}
