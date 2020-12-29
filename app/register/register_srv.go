package register

import (
	"github.com/donnol/jdnote/app"

	"github.com/donnol/jdnote/services/authsrv"
	"github.com/donnol/jdnote/services/filesrv"
	"github.com/donnol/jdnote/services/notesrv"
	"github.com/donnol/jdnote/services/timesrv"
	"github.com/donnol/jdnote/services/usersrv"
)

func registerSrvProvider(appObj *app.App) {
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

		// more...
	)
}
