package rolemodel

import (
	"github.com/donnol/jdnote/utils/context"
)

type IRole interface {
	Add(ctx context.Context, e Entity) (id int, err error)
	GetByID(ctx context.Context, id int) (e Entity, err error)
}

func New() IRole {
	return &roleImpl{}
}
