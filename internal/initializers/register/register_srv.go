package register

import (
	"github.com/donnol/jdnote/internal/initializers"

	"github.com/donnol/jdnote/services/authsrv"
	"github.com/donnol/jdnote/services/filesrv"
	"github.com/donnol/jdnote/services/notesrv"
	"github.com/donnol/jdnote/services/timesrv"
	"github.com/donnol/jdnote/services/usersrv"
)

func registerSrvProvider(appObj *initializers.App) {
	appObj.MustRegisterProvider(
		initializers.ProviderOption{
			Provider: usersrv.New,
			Mock:     &usersrv.UserMock{},
		},
		initializers.ProviderOption{
			Provider: authsrv.New,
			Mock:     &authsrv.AuthMock{},
		},
		initializers.ProviderOption{
			Provider: notesrv.New,
			Mock:     &notesrv.NoteMock{},
		},
		initializers.ProviderOption{
			Provider: timesrv.New,
			Mock:     &timesrv.TimeMock{},
		},
		initializers.ProviderOption{
			Provider: filesrv.NewIFile,
			Mock:     &filesrv.FileMock{},
		},

		// more...
	)
}
