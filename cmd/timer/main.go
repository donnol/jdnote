package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/donnol/jdnote/internal/initializers"
	"github.com/donnol/jdnote/internal/initializers/register"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	appObj, cctx := initializers.New(ctx)
	defer appObj.Cancel()
	logger := appObj.Logger()

	// 先注册ArounderMap，因为后面的依赖注入和proxy依赖它
	appObj.RegisterArounderMap(register.GetArounder())

	// 注入依赖并注册定时器handler
	register.InjectAndRegisterTimerHandler(cctx, appObj)

	// 监听终止信号
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)

		sig := <-sigint
		logger.Debugf("Recv interrupt signal, %v", sig)

		cronCtx := appObj.StopTimer()
		// 收到终止信号后，最多再等n秒
		timer := time.NewTimer(5 * 60 * time.Second)
		select {
		case <-cronCtx.Done():
			logger.Infof("timer stop.\n")
		case t := <-timer.C:
			logger.Infof("timer exceed, %v.\n", t)
		}

		// 关闭管道，让进程能顺利停止
		close(idleConnsClosed)
	}()

	// 启动
	if err := appObj.RunTimer(); err != nil {
		logger.Errorf("Timer err: %+v\n", err)
	}

	// 放在最后，确保前面的工作已完成
	<-idleConnsClosed

	logger.Infof("All finish.\n")
}
