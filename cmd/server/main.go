package main

import (
	_ "github.com/donnol/jdnote/api/auth"
	"github.com/donnol/jdnote/route"
	utillog "github.com/donnol/jdnote/utils/log"
)

func main() {
	utillog.Debugf("jdnote server start.")

	if err := route.DefaultRouter.Run(":8810"); err != nil {
		panic(err)
	}
}
