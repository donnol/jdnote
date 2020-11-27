package roleactionstore

import (
	"github.com/donnol/jdnote/models/roleactionmodel"
	"github.com/donnol/jdnote/utils/context"
)

type IRoleAction interface {
	Add(ctx context.Context, e roleactionmodel.Entity) (id int, err error)
	CheckPerm(ctx context.Context, perms []string) error
}

func New() IRoleAction {
	return &roleActionImpl{}
}
