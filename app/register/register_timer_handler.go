package register

import (
	"github.com/donnol/jdnote/app"

	"github.com/donnol/jdnote/services/notesrv"

	"github.com/donnol/jdnote/utils/context"
)

// TODO: 根据timerHandler里的字段类型所关联方法里返回值类型是timer.Job/timer.FuncJob来决定是否注册
type timerHandler struct {
	noteSrv notesrv.INote

	// more...
}

func InjectAndRegisterTimerHandler(cctx context.Context, appObj *app.App) {
	// 注册provider
	registerProvider(appObj)

	th := &timerHandler{}
	appObj.MustInject(th)

	spec := "* * * * *"
	appObj.RegisterTimerHandler(spec, th.noteSrv.Timer(cctx))

	// more...
}
