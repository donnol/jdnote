package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/donnol/jdnote/app"
	utillog "github.com/donnol/tools/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "net/http/pprof"

	// 注入路由
	"github.com/donnol/jdnote/api/auth"
	"github.com/donnol/jdnote/api/file"
	"github.com/donnol/jdnote/api/note"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	appObj, cctx := app.New(ctx)

	// 注册路由
	appObj.Register(cctx, &auth.Auth{})
	appObj.Register(cctx, &file.File{})
	appObj.Register(cctx, &note.Note{})
	router := appObj.Router()

	// 静态文件
	router.StaticFS("/static", http.Dir("dist"))

	// 服务器
	port := fmt.Sprintf(":%d", appObj.GetConfig().Server.Port)
	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	// 启动pprof
	go func() {
		utillog.Debugf("Pprof server start\n")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// 启动prometheus
	go func() {
		utillog.Debugf("Prometheus server start\n")
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe("localhost:6660", nil))
	}()

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
	utillog.Debugf("Server start at %v. Listening '%s'", time.Now().Format("2006-01-02 15:04:05"), port)
	if err := srv.ListenAndServe(); err != nil {
		utillog.Debugf("HTTP server ListenAndServe: %v", err)
	}

	// 放在最后，确保前面的工作已完成
	<-idleConnsClosed
}
