package rolestore

import (
	"context"

	"github.com/donnol/jdnote/models/rolemodel"
	"github.com/donnol/jdnote/utils/store/db"
)

type IRole interface {
	Add(ctx context.Context, e rolemodel.Entity) (id int, err error)
	GetByID(ctx context.Context, id int) (e rolemodel.Entity, err error)
}

func New(
	db db.DB,
) IRole {
	return &roleImpl{
		db: db,
	}
}
