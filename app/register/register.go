package register

import (
	"github.com/donnol/jdnote/app"
	"github.com/donnol/jdnote/stores/actionstore"
	"github.com/donnol/jdnote/stores/filestore"
	"github.com/donnol/jdnote/stores/notestore"
	"github.com/donnol/jdnote/stores/roleactionstore"
	"github.com/donnol/jdnote/stores/rolestore"
	"github.com/donnol/jdnote/stores/userrolestore"
	"github.com/donnol/jdnote/stores/userstore"

	"github.com/donnol/jdnote/api/authapi"
	"github.com/donnol/jdnote/api/fileapi"
	"github.com/donnol/jdnote/api/noteapi"
	"github.com/donnol/jdnote/api/timeapi"

	"github.com/donnol/jdnote/services/authsrv"
	"github.com/donnol/jdnote/services/filesrv"
	"github.com/donnol/jdnote/services/notesrv"
	"github.com/donnol/jdnote/services/timesrv"
	"github.com/donnol/jdnote/services/usersrv"

	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/queue"

	"github.com/donnol/tools/inject"
	"github.com/donnol/tools/log"
)

// RegisterAll 注册所有模块
func RegisterAll(cctx context.Context, appObj *app.App) {
	logger := appObj.Logger()
	trigger := appObj.Trigger()

	// 注入通用provider
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
	// store
	appObj.MustRegisterProvider(
		app.ProviderOption{
			Provider: userstore.New,
			Mock:     &userstore.UserMock{},
			Hooks: []inject.Hook{
				&app.TimeHook{},
			},
		},
		app.ProviderOption{
			Provider: userrolestore.New,
		},
		app.ProviderOption{
			Provider: rolestore.New,
		},
		app.ProviderOption{
			Provider: actionstore.New,
		},
		app.ProviderOption{
			Provider: roleactionstore.New,
		},
		app.ProviderOption{
			Provider: notestore.New,
		},
		app.ProviderOption{
			Provider: filestore.NewIFile,
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
		app.ProviderOption{
			Provider: filesrv.NewIFile,
		},
	)

	// 注入依赖，并注册路由
	appObj.Register(cctx, &authapi.Auth{})
	appObj.Register(cctx, &fileapi.File{})
	appObj.Register(cctx, &noteapi.Note{})
	appObj.Register(cctx, &noteapi.Front{})
	appObj.Register(cctx, &timeapi.Time{})
}
