package models

import (
	"github.com/donnol/jdnote/utils/context"
	utillog "github.com/donnol/tools/log"
)

// DefaultCtx 默认
func DefaultCtx() context.Context {
	return context.New(defaultDB, utillog.Default(), 0)
}
