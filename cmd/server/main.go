package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/donnol/jdnote/app"
	"github.com/donnol/jdnote/app/register"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	appObj, cctx := app.New(ctx)
	defer appObj.Cancel()
	logger := appObj.Logger()

	// 先注册ArounderMap，因为后面的依赖注入和proxy依赖它
	appObj.RegisterArounderMap(register.GetArounder())

	// 注册
	register.InjectAndRegisterRouter(cctx, appObj)

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
