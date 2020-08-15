package main

import (
	"time"

	utillog "github.com/donnol/tools/log"
	"github.com/robfig/cron"
)

func main() {
	obj := cron.New()
	_ = obj
	utillog.Debugf("Start timer: %v\n", time.Now())

	// TODO:

	// 保持运行
	for {

	}
}
