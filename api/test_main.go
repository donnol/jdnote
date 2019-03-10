package api

import (
	"sync"

	"github.com/donnol/jdnote/route"
)

// TestPort 测试端口
const TestPort = 8820

var startOnce = new(sync.Once)

// TestMain 测试初始化
func TestMain() {
	// 开启服务
	startOnce.Do(func() {
		go func() {
			if err := route.StartServer(TestPort); err != nil {
				panic(err)
			}
		}()
	})
}
