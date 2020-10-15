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
	"github.com/donnol/tools/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "net/http/pprof"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	appObj, cctx := app.New(ctx)
	defer appObj.Cancel()

	// 注入provider
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

	// 启动pprof
	go func() {
		log.Debugf("Pprof server start\n")
		log.Errorf("pprof ListenAndServe err: %+v\n", http.ListenAndServe("localhost:6060", nil))
	}()

	// 启动prometheus
	go func() {
		log.Debugf("Prometheus server start\n")
		http.Handle("/metrics", promhttp.Handler())
		log.Errorf("prometheus ListenAndServe err: %+v\n", http.ListenAndServe("localhost:6660", nil))
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
			log.Debugf("Recv interrupt signal.")
		case <-sigterm:
			log.Debugf("Recv terminal signal.")
		}

		if err := appObj.ShutdownServer(ctx); err != nil {
			log.Errorf("ShutdownServer failed: %+v\n", err)
		}

		// 关闭管道，让进程能顺利停止
		close(idleConnsClosed)
	}()

	// 开启服务器
	if err := appObj.Run(); err != nil {
		log.Errorf("Server err: %+v\n", err)
	}

	// 放在最后，确保前面的工作已完成
	<-idleConnsClosed
}
