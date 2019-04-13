package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/donnol/jdnote/route"
	utillog "github.com/donnol/jdnote/utils/log"

	// 注入路由
	_ "github.com/donnol/jdnote/api/auth"
)

func main() {
	// 配置服务器
	port := ":8810"
	srv := &http.Server{
		Addr:    port,
		Handler: route.DefaultRouter,
	}

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

		// 优雅关闭
		if err := srv.Shutdown(context.Background()); err != nil {
			utillog.Debugf("HTTP server Shutdown: %v", err)
		}

		// 关闭管道，让进程能顺利停止
		close(idleConnsClosed)
	}()

	// 开启服务器
	utillog.Debugf("Server start. Listen '%s'", port)
	if err := srv.ListenAndServe(); err != nil {
		utillog.Debugf("HTTP server ListenAndServe: %v", err)
	}

	// 放在最后，确保前面的工作已完成
	<-idleConnsClosed
}
