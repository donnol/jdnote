package app

import (
	"github.com/donnol/tools/inject"
)

// 内置ioc
var (
	defaultIoc = inject.NewIoc(true)
)

func RegisterProvider(v interface{}) (err error) {
	return defaultIoc.RegisterProvider(v)
}

func Inject(v interface{}) (err error) {
	return defaultIoc.Inject(v)
}

func MustRegisterProvider(v interface{}) {
	if err := defaultIoc.RegisterProvider(v); err != nil {
		panic(err)
	}
}

func MustInject(v interface{}) {
	if err := defaultIoc.Inject(v); err != nil {
		panic(err)
	}
}
