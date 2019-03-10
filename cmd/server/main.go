package main

import (
	_ "github.com/donnol/jdnote/api/auth"
	"github.com/donnol/jdnote/route"
	utillog "github.com/donnol/jdnote/utils/log"
)

func main() {
	utillog.Debugf("jdnote server start.")

	port := 8810
	if err := route.StartServer(port); err != nil {
		panic(err)
	}
}
