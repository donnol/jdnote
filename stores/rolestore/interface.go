package rolestore

import (
	"github.com/donnol/jdnote/models/rolemodel"
	"github.com/donnol/jdnote/utils/context"
)

type IRole interface {
	Add(ctx context.Context, e rolemodel.Entity) (id int, err error)
	GetByID(ctx context.Context, id int) (e rolemodel.Entity, err error)
}

func New() IRole {
	return &roleImpl{}
}
