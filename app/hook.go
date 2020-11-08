package app

import (
	"time"

	"github.com/donnol/tools/inject"
)

type TimeHook struct {
	begin time.Time
	end   time.Time
}

func (hook *TimeHook) Before(pctx inject.ProxyContext) {
	hook.begin = time.Now()
	pctx.Logf("begin: %v\n", hook.begin)
}

func (hook *TimeHook) After(pctx inject.ProxyContext) {
	hook.end = time.Now()
	pctx.Logf("end: %v\n", hook.end)

	// 耗时
	used := hook.end.Sub(hook.begin)
	pctx.Logf("used: %v\n", used)
}
