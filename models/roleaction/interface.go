package roleaction

import (
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/inject"
)

type IRoleAction interface {
	Add(ctx context.Context, e Entity) (id int, err error)
	CheckPerm(ctx context.Context, perms []string) error
}

func New() IRoleAction {
	return &roleActionImpl{}
}

func init() {
	inject.MustRegisterProvider(New)
}
