package register

import (
	"github.com/donnol/jdnote/app"
)

func registerProvider(appObj *app.App) {
	// common
	registerCommonProvider(appObj)

	// store
	registerStoreProvider(appObj)

	// service
	registerSrvProvider(appObj)
}
