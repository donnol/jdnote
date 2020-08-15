package userrole

import (
	"github.com/donnol/jdnote/app"
	"github.com/donnol/jdnote/utils/context"
)

type IUserRole interface {
	Add(ctx context.Context, e Entity) (id int, err error)
	GetByUserID(ctx context.Context, userID int) (list []Entity, err error)
}

func New() IUserRole {
	return &userRoleImpl{}
}

func init() {
	app.MustRegisterProvider(New)
}
