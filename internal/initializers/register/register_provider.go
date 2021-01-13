package register

import (
	"github.com/donnol/jdnote/internal/initializers"
)

func registerProvider(appObj *initializers.App) {
	// common
	registerCommonProvider(appObj)

	// store
	registerStoreProvider(appObj)

	// service
	registerSrvProvider(appObj)
}
