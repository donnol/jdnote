package userrolestore

import (
	"context"

	"github.com/donnol/jdnote/models/userrolemodel"
	"github.com/donnol/jdnote/utils/store/db"
)

type IUserRole interface {
	Add(ctx context.Context, e userrolemodel.Entity) (id int, err error)
	GetByUserID(ctx context.Context, userID int) (list []userrolemodel.Entity, err error)
}

func New(
	db db.DB,
) IUserRole {
	return &userRoleImpl{
		db: db,
	}
}
