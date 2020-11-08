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

	"github.com/donnol/tools/inject"
	"github.com/donnol/tools/log"
)

func appRegister(cctx context.Context, appObj *app.App) {
	logger := appObj.Logger()
	trigger := appObj.Trigger()

	// 注入provider
	appObj.MustRegisterProvider(
		app.ProviderOption{
			Provider: func() log.Logger {
				return logger
			},
		},
		app.ProviderOption{
			Provider: func() queue.Trigger {
				return trigger
			},
		},
	)
	// model
	appObj.MustRegisterProvider(
		app.ProviderOption{
			Provider: usermodel.New,
			Mock:     &usermodel.UserMock{},
			Hooks: []inject.Hook{
				&app.TimeHook{},
			},
		},
		app.ProviderOption{
			Provider: userrolemodel.New,
		},
		app.ProviderOption{
			Provider: rolemodel.New,
		},
		app.ProviderOption{
			Provider: actionmodel.New,
		},
		app.ProviderOption{
			Provider: roleactionmodel.New,
		},
		app.ProviderOption{
			Provider: notemodel.New,
		},
	)
	// service
	appObj.MustRegisterProvider(
		app.ProviderOption{
			Provider: usersrv.New,
		},
		app.ProviderOption{
			Provider: authsrv.New,
		},
		app.ProviderOption{
			Provider: notesrv.New,
		},
		app.ProviderOption{
			Provider: timesrv.New,
		},
	)

	// 注入依赖，并注册路由
	appObj.Register(cctx, &authapi.Auth{})
	appObj.Register(cctx, &fileapi.File{})
	appObj.Register(cctx, &noteapi.Note{})
	appObj.Register(cctx, &timeapi.Time{})
}
