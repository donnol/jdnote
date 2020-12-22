package register

import (
	"github.com/donnol/jdnote/stores/notestore"
	"github.com/donnol/jdnote/stores/userstore"

	"github.com/donnol/tools/inject"
)

// GetArounder 获取特定Arounder配置
// pctx需要执行的特别函数
// 在这里，可以对method, args, result做手脚，但是均不建议这样做
// 推荐做法是只在方法调用前后做一些操作
func GetArounder() inject.ArounderMap {
	m := make(inject.ArounderMap)

	m = m.Merge(userstore.GetArounder())
	m = m.Merge(notestore.GetArounder())

	// more...

	return m
}
