package register

import (
	"reflect"

	"github.com/donnol/jdnote/stores/notestore"
	"github.com/donnol/jdnote/stores/userstore"

	"github.com/donnol/tools/inject"
)

var (
	// pctx需要执行的特别函数
	// 在这里，可以对method, args, result做手脚，但是均不建议这样做
	// 所以最好是只在方法调用前后做一些操作
	ArounderMap = map[inject.ProxyContext]inject.AroundFunc{
		userstore.UserMockGetByIDProxyContext: func(pctx inject.ProxyContext, method reflect.Value, args []reflect.Value) []reflect.Value {
			var result []reflect.Value

			pctx.Logf("do some specify things before\n")

			result = method.Call(args)

			pctx.Logf("do some specify things after\n")

			return result
		},
		notestore.NoterMockAddOneProxyContext: func(pctx inject.ProxyContext, method reflect.Value, args []reflect.Value) []reflect.Value {
			var result []reflect.Value

			// 打印参数
			for _, arg := range args {
				if arg.CanInterface() {
					pctx.Logf("arg: %v, %s\n", arg.Interface(), arg.String())
				} else {
					pctx.Logf("arg: %v, %s\n", arg, arg.String())
				}
			}

			result = method.Call(args)

			// 打印结果
			for _, r := range result {
				if r.CanInterface() {
					pctx.Logf("result: %v, %s\n", r.Interface(), r.String())
				} else {
					pctx.Logf("result: %v, %s\n", r, r.String())
				}
			}

			return result
		},
		// more...
	}
)
