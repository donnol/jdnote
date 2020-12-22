package userstore

import (
	"reflect"

	"github.com/donnol/tools/inject"
)

func GetArounder() inject.ArounderMap {
	return inject.ArounderMap{
		UserMockGetByIDProxyContext: func(pctx inject.ProxyContext, method reflect.Value, args []reflect.Value) []reflect.Value {
			var result []reflect.Value

			pctx.Logf("do some specify things before\n")

			result = method.Call(args)

			pctx.Logf("do some specify things after\n")

			return result
		},
	}
}
