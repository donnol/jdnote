package roleaction

import (
	"github.com/donnol/jdnote/app"
	"github.com/donnol/jdnote/utils/context"
)

type IRoleAction interface {
	Add(ctx context.Context, e Entity) (id int, err error)
	CheckPerm(ctx context.Context, perms []string) error
}

func New() IRoleAction {
	return &roleActionImpl{}
}

func init() {
	app.MustRegisterProvider(New)
}
