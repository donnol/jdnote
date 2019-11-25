package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/donnol/jdnote/config"
	utillog "github.com/donnol/jdnote/utils/log"
	"github.com/google/gops/agent"
)

func main() {
	addr := fmt.Sprintf(":%d", config.Default().Server.Port+1)

	// 监听终止信号
	idleConnsClosed := make(chan struct{})
	go func() {
		// ctrl+c停止
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)

		// docker stop会发这个信号给进程
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGTERM)

		select {
		case <-sigint:
			utillog.Debugf("Recv interrupt signal.")
		case <-sigterm:
			utillog.Debugf("Recv terminal signal.")
		}

		// 关闭
		agent.Close()

		// 关闭管道，让进程能顺利停止
		close(idleConnsClosed)
	}()

	utillog.Debugf("Gops start at %v. Listening '%s'", time.Now().Format("2006-01-02 15:04:05"), addr)
	if err := agent.Listen(agent.Options{
		Addr:            addr,
		ShutdownCleanup: true,
	}); err != nil {
		log.Fatal(err)
	}

	// 放在最后，确保前面的工作已完成
	<-idleConnsClosed
}
