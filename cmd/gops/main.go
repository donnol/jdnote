package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	utillog "github.com/donnol/tools/log"
	"github.com/google/gops/agent"
)

func main() {
	addr := fmt.Sprintf(":%d", 9091)

	// 监听终止信号
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)

		sig := <-sigint
		utillog.Debugf("Recv signal: %v.", sig)

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
