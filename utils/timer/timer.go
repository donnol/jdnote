package timer

import (
	"time"

	"github.com/donnol/tools/log"
	"github.com/robfig/cron/v3"
)

type Job = cron.Job

type FuncJob = cron.FuncJob

func WithTimeConsuming(mark string, f FuncJob) FuncJob {
	return func() {
		begin := time.Now()
		defer func() {
			log.Default().Infof(mark+"used time: %v\n", time.Since(begin))
		}()

		f()
	}
}
