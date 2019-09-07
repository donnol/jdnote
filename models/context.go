package model

import (
	"log"
	"os"

	"github.com/donnol/jdnote/utils/context"
	utillog "github.com/donnol/jdnote/utils/log"
)

// DefaultCtx 默认
func DefaultCtx() context.Context {
	return context.New(defaultDB, utillog.New(os.Stdout, "", log.LstdFlags), 0)
}
