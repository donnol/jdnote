package initializers

import "time"

// Option 选项，控制是否需要初始化等
type Option struct {
	timeout time.Duration
}

// OptionSetter 给Option赋值
type OptionSetter func(*Option)

// SetTimeout 因为Option的字段都是非导出的，需要提供方法给用户来获取相应字段的OptionSetter函数
// 拿到setter函数之后，再将它传给New方法就好了
func SetTimeout(timeout time.Duration) OptionSetter {
	return func(opt *Option) {
		opt.timeout = timeout
	}
}

// checkRequire 检查必填项，如果没设置，或报错，或使用默认值
func (opt *Option) checkRequire() error {
	// 没有则设置默认值
	if opt.timeout == 0 {
		opt.timeout = time.Hour * 1
	}

	// 没有则报错

	return nil
}
