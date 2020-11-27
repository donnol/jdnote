package actionstore

import (
	"github.com/donnol/jdnote/models/actionmodel"
	"github.com/donnol/jdnote/utils/context"
)

type IAction interface {
	Add(ctx context.Context, e actionmodel.Entity) (id int, err error)
	GetByID(ctx context.Context, id int) (e actionmodel.Entity, err error)
}

func New() IAction {
	return &actionImpl{}
}
