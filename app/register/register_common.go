package register

import (
	"github.com/donnol/jdnote/app"

	"github.com/donnol/jdnote/utils/queue"

	"github.com/donnol/tools/log"
)

func registerCommonProvider(appObj *app.App) {
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

		// more...
	)
}
