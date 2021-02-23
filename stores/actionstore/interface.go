package actionstore

import (
	"context"

	"github.com/donnol/jdnote/models/actionmodel"
	"github.com/donnol/jdnote/utils/store/db"
)

type IAction interface {
	Add(ctx context.Context, e actionmodel.Entity) (id int, err error)
	GetByID(ctx context.Context, id int) (e actionmodel.Entity, err error)
}

func New(
	db db.DB,
) IAction {
	return &actionImpl{
		db: db,
	}
}
