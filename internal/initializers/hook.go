package initializers

import (
	"reflect"
	"time"

	"github.com/donnol/tools/inject"
)

func GetArounder(arounderMap map[inject.ProxyContext]inject.AroundFunc) inject.AroundFunc {
	return func(pctx inject.ProxyContext, method reflect.Value, args []reflect.Value) []reflect.Value {
		var result []reflect.Value

		// 执行前
		begin := time.Now()

		// 针对pctx的操作
		around, ok := arounderMap[pctx]
		if ok {
			result = around(pctx, method, args)
		} else {
			result = method.Call(args)
		}

		// 执行后
		used := time.Since(begin)
		processUsedTime(pctx, used)

		return result
	}
}

func processUsedTime(pctx inject.ProxyContext, used time.Duration) {
	usedSec := used.Seconds()
	switch {
	case usedSec >= 10:
		pctx.Logf("| Used Time | terrible long: %v\n", used)
	case usedSec >= 5:
		pctx.Logf("| Used Time | too long: %v\n", used)
	case usedSec >= 3:
		pctx.Logf("| Used Time | very long: %v\n", used)
	case usedSec >= 1:
		pctx.Logf("| Used Time | a bit long: %v\n", used)
	case usedSec >= 0.5:
		pctx.Logf("| Used Time | normal: %v\n", used)
	default:
		pctx.Logf("| Used Time | good: %v\n", used)
	}
}
