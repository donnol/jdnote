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
		},
		app.ProviderOption{
			Provider: userrolestore.New,
			Mock:     &userrolestore.UserRoleMock{},
		},
		app.ProviderOption{
			Provider: rolestore.New,
			Mock:     &rolestore.RoleMock{},
		},
		app.ProviderOption{
			Provider: actionstore.New,
			Mock:     &actionstore.ActionMock{},
		},
		app.ProviderOption{
			Provider: roleactionstore.New,
			Mock:     &roleactionstore.RoleActionMock{},
		},
		app.ProviderOption{
			Provider: notestore.New,
			Mock:     &notestore.NoterMock{},
		},
		app.ProviderOption{
			Provider: filestore.NewIFile,
			Mock:     &filestore.FileMock{},
		},
	)
	// service
	appObj.MustRegisterProvider(
		app.ProviderOption{
			Provider: usersrv.New,
			Mock:     &usersrv.UserMock{},
		},
		app.ProviderOption{
			Provider: authsrv.New,
			Mock:     &authsrv.AuthMock{},
		},
		app.ProviderOption{
			Provider: notesrv.New,
			Mock:     &notesrv.NoteMock{},
		},
		app.ProviderOption{
			Provider: timesrv.New,
			Mock:     &timesrv.TimeMock{},
		},
		app.ProviderOption{
			Provider: filesrv.NewIFile,
			Mock:     &filesrv.FileMock{},
		},
	)

	// 注入依赖，并注册路由
	appObj.Register(cctx, &authapi.Auth{})
	appObj.Register(cctx, &fileapi.File{})
	appObj.Register(cctx, &noteapi.Note{})
	appObj.Register(cctx, &noteapi.Front{})
	appObj.Register(cctx, &timeapi.Time{})
}
