package notestore

import (
	"reflect"

	"github.com/donnol/jdnote/models/notemodel"
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/tools/inject"
)

func GetArounder() inject.ArounderMap {
	return inject.ArounderMap{
		NoterMockAddOneProxyContext: func(pctx inject.ProxyContext, method reflect.Value, args []reflect.Value) []reflect.Value {
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
		NoterMockModProxyContext: func(pctx inject.ProxyContext, method reflect.Value, args []reflect.Value) []reflect.Value {
			var result []reflect.Value

			for i, arg := range args {
				if arg.CanInterface() {
					argType := arg.Type()
					pctx.Logf("arg: %v, %s\n", arg.Interface(), arg.String())

					// 查看ctx信息
					if i == 0 {
						ctx, ok := arg.Interface().(context.Context)
						if ok {
							values, err := context.GetAllValue(ctx)
							if err != nil {
								pctx.Logf("GetAllValue err: %+v\n", err)
							}
							pctx.Logf("GetAllValue: %+v\n", values)
						}
					}

					// 以参数位置和参数类型判断
					// 类型是 *notemodel.Entity
					paramPtrType := reflect.TypeOf((*notemodel.Entity)(nil))
					if i == 2 && argType == paramPtrType { // 如果改了方法签名（加减参数，改变类型），这里就对不上了，那么，是不是把这个方法放到离原方法更近的地方好呢？
						// 修改参数里的标题
						// reflectx.SetStructFieldValue(arg, "Title", arg.Interface().(*notemodel.Entity).Title+"_hahah")

						pctx.Logf("arg notemodel.Entity: %v, %s\n", arg.Interface(), arg.String())
					}
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
	}
}
