package flow

import (
	"context"
)

// Register 注册
func Register(Flow, []Hook) {} // 注册新类型，并绑定钩子

// Flow 流
type Flow interface {
	// 准备
	Prepare(context.Context) error
	// 执行
	Do(context.Context) error
	// 完成
	Finish(context.Context) error
	// 收尾
	Cleanup(context.Context) error
}

// Timing 时机
type Timing int

// 时机枚举
const (
	Before Timing = iota + 1
	After
)

// Hook 钩子
type Hook interface {
	When() Timing
	Do(context.Context) error
}
