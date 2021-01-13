package register

import (
	"github.com/donnol/jdnote/internal/initializers"

	"github.com/donnol/jdnote/stores/actionstore"
	"github.com/donnol/jdnote/stores/filestore"
	"github.com/donnol/jdnote/stores/notestore"
	"github.com/donnol/jdnote/stores/roleactionstore"
	"github.com/donnol/jdnote/stores/rolestore"
	"github.com/donnol/jdnote/stores/userrolestore"
	"github.com/donnol/jdnote/stores/userstore"
)

func registerStoreProvider(appObj *initializers.App) {
	appObj.MustRegisterProvider(
		initializers.ProviderOption{
			Provider: userstore.New,
			Mock:     &userstore.UserMock{},
		},
		initializers.ProviderOption{
			Provider: userrolestore.New,
			Mock:     &userrolestore.UserRoleMock{},
		},
		initializers.ProviderOption{
			Provider: rolestore.New,
			Mock:     &rolestore.RoleMock{},
		},
		initializers.ProviderOption{
			Provider: actionstore.New,
			Mock:     &actionstore.ActionMock{},
		},
		initializers.ProviderOption{
			Provider: roleactionstore.New,
			Mock:     &roleactionstore.RoleActionMock{},
		},
		initializers.ProviderOption{
			Provider: notestore.New,
			Mock:     &notestore.NoterMock{},
		},
		initializers.ProviderOption{
			Provider: filestore.NewIFile,
			Mock:     &filestore.FileMock{},
		},

		// more...
	)
}
