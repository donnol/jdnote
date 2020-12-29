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
)

func registerStoreProvider(appObj *app.App) {
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

		// more...
	)
}
