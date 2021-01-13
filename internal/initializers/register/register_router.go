package register

import (
	"github.com/donnol/jdnote/api/authapi"
	"github.com/donnol/jdnote/api/fileapi"
	"github.com/donnol/jdnote/api/noteapi"
	"github.com/donnol/jdnote/api/timeapi"
	"github.com/donnol/jdnote/internal/initializers"

	"github.com/donnol/jdnote/utils/context"
)

// InjectAndRegisterRouter 注入依赖并注册路由
func InjectAndRegisterRouter(cctx context.Context, appObj *initializers.App) {
	// 注册provider
	registerProvider(appObj)

	// 注入依赖，并注册路由
	for _, target := range getRouterTargets() {
		appObj.RegisterRouterWithInject(cctx, target)
	}
}

func getRouterTargets() []interface{} {
	return []interface{}{
		&authapi.Auth{},
		&fileapi.File{},
		&noteapi.Note{},
		&noteapi.Front{},
		&timeapi.Time{},

		// more...
	}
}
