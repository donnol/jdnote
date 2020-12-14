package register

import (
	"reflect"

	"github.com/donnol/jdnote/stores/userstore"

	"github.com/donnol/tools/inject"
)

var (
	// pctx需要执行的特别函数
	ArounderMap = map[inject.ProxyContext]inject.AroundFunc{
		userstore.UserMockGetByIDProxyContext: func(pctx inject.ProxyContext, method reflect.Value, args []reflect.Value) []reflect.Value {
			var result []reflect.Value

			pctx.Logf("do some specify things before\n")

			result = method.Call(args)

			pctx.Logf("do some specify things after\n")

			return result
		},
		// more...
	}
)
