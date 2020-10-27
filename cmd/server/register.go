package main

import (
	"github.com/donnol/jdnote/app"

	"github.com/donnol/jdnote/api/authapi"
	"github.com/donnol/jdnote/api/fileapi"
	"github.com/donnol/jdnote/api/noteapi"
	"github.com/donnol/jdnote/api/timeapi"

	"github.com/donnol/jdnote/models/actionmodel"
	"github.com/donnol/jdnote/models/notemodel"
	"github.com/donnol/jdnote/models/roleactionmodel"
	"github.com/donnol/jdnote/models/rolemodel"
	"github.com/donnol/jdnote/models/usermodel"
	"github.com/donnol/jdnote/models/userrolemodel"

	"github.com/donnol/jdnote/services/authsrv"
	"github.com/donnol/jdnote/services/notesrv"
	"github.com/donnol/jdnote/services/timesrv"
	"github.com/donnol/jdnote/services/usersrv"

	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/queue"

	"github.com/donnol/tools/log"
)

func appRegister(cctx context.Context, appObj *app.App) {
	logger := appObj.Logger()
	trigger := appObj.Trigger()

	// 注入provider
	appObj.MustRegisterProvider(
		func() log.Logger {
			return logger
		},
		func() queue.Trigger {
			return trigger
		},
	)
	// model
	appObj.MustRegisterProvider(
		usermodel.New,
		userrolemodel.New,
		rolemodel.New,
		actionmodel.New,
		roleactionmodel.New,
		notemodel.New,
	)
	// service
	appObj.MustRegisterProvider(
		usersrv.New,
		authsrv.New,
		notesrv.New,
		timesrv.New,
	)

	// 注入依赖，并注册路由
	appObj.Register(cctx, &authapi.Auth{})
	appObj.Register(cctx, &fileapi.File{})
	appObj.Register(cctx, &noteapi.Note{})
	appObj.Register(cctx, &timeapi.Time{})
}
