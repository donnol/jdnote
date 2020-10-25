package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/donnol/jdnote/api/authapi"
	"github.com/donnol/jdnote/api/fileapi"
	"github.com/donnol/jdnote/api/noteapi"
	"github.com/donnol/jdnote/models/actionmodel"
	"github.com/donnol/jdnote/models/notemodel"
	"github.com/donnol/jdnote/models/roleactionmodel"
	"github.com/donnol/jdnote/models/rolemodel"
	"github.com/donnol/jdnote/models/usermodel"
	"github.com/donnol/jdnote/models/userrolemodel"
	"github.com/donnol/jdnote/services/authsrv"
	"github.com/donnol/jdnote/services/notesrv"
	"github.com/donnol/jdnote/services/usersrv"

	"github.com/donnol/jdnote/app"
	"github.com/donnol/jdnote/utils/queue"

	"github.com/donnol/tools/log"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	appObj, cctx := app.New(ctx)
	defer appObj.Cancel()
	logger := appObj.Logger()
	trigger := appObj.Trigger()

	// 注入provider
	appObj.MustRegisterProvider(
		func() log.Logger {
			return logger
		},
		func() queue.Trigger {
			return trigger
		},
	)
	// model
	appObj.MustRegisterProvider(
		usermodel.New,
		userrolemodel.New,
		rolemodel.New,
		actionmodel.New,
		roleactionmodel.New,
		notemodel.New,
	)
	// service
	appObj.MustRegisterProvider(
		usersrv.New,
		authsrv.New,
		notesrv.New,
	)

	// 注入依赖，并注册路由
	appObj.Register(cctx, &authapi.Auth{})
	appObj.Register(cctx, &fileapi.File{})
	appObj.Register(cctx, &noteapi.Note{})

	// 静态文件
	appObj.StaticFS("/static", http.Dir("dist"))

	// 监听终止信号
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)

		sig := <-sigint
		logger.Debugf("Recv interrupt signal, %v", sig)

		if err := appObj.ShutdownServer(ctx); err != nil {
			logger.Errorf("ShutdownServer failed: %+v\n", err)
		}

		// 关闭管道，让进程能顺利停止
		close(idleConnsClosed)
	}()

	if err := appObj.RunPprof(); err != nil {
		logger.Errorf("Pprof err: %+v\n", err)
	}

	if err := appObj.RunPrometheus(); err != nil {
		logger.Errorf("Prometheus err: %+v\n", err)
	}

	// 开启服务器
	if err := appObj.Run(); err != nil {
		logger.Errorf("Server err: %+v\n", err)
	}

	logger.Infof("All finish.\n")

	// 放在最后，确保前面的工作已完成
	<-idleConnsClosed
}
