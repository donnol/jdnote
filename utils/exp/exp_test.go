package exp

import (
	"log"
	"testing"
)

func TestSrv(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// init
	Srv := &Srv{}
	log.Printf("%p\n", Srv)
	log.Printf("%p\n", Srv.Base)
	newDB := &db{}
	log.Printf("%p\n", newDB)
	Srv.DB = newDB
	log.Printf("%p\n", Srv.Base)
	t.Logf("%+v\n", Srv)
	Srv.Init()

	// get
	r, err := Srv.DB.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", r)

	Srv.Begin()
	t.Logf("%+v\n", Srv)
}
