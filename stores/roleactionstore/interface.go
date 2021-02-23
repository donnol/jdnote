package roleactionstore

import (
	"context"

	"github.com/donnol/jdnote/models/roleactionmodel"
	"github.com/donnol/jdnote/utils/store/db"
)

type IRoleAction interface {
	Add(ctx context.Context, e roleactionmodel.Entity) (id int, err error)
	CheckPerm(ctx context.Context, perms []string) error
}

func New(
	db db.DB,
) IRoleAction {
	return &roleActionImpl{
		db: db,
	}
}
