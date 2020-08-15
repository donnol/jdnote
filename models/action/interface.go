package action

import (
	"github.com/donnol/jdnote/app"
	"github.com/donnol/jdnote/utils/context"
)

type IAction interface {
	Add(ctx context.Context, e Entity) (id int, err error)
	GetByID(ctx context.Context, id int) (e Entity, err error)
}

func New() IAction {
	return &actionImpl{}
}

func init() {
	app.MustRegisterProvider(New)
}
