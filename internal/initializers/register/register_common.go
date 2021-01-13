package register

import (
	"github.com/donnol/jdnote/internal/initializers"

	"github.com/donnol/jdnote/utils/queue"

	"github.com/donnol/tools/log"
)

func registerCommonProvider(appObj *initializers.App) {
	logger := appObj.Logger()
	trigger := appObj.Trigger()

	// 注入通用provider
	appObj.MustRegisterProvider(
		initializers.ProviderOption{
			Provider: func() log.Logger {
				return logger
			},
		},
		initializers.ProviderOption{
			Provider: func() queue.Trigger {
				return trigger
			},
		},

		// more...
	)
}
