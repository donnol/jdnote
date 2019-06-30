package exp

import (
	"log"
	"testing"
)

func TestAo(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// init
	ao := &Ao{}
	log.Printf("%p\n", ao)
	log.Printf("%p\n", ao.Base)
	newDB := &db{}
	log.Printf("%p\n", newDB)
	ao.DB = newDB
	log.Printf("%p\n", ao.Base)
	t.Logf("%+v\n", ao)
	ao.Init()

	// get
	r, err := ao.DB.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", r)

	ao.Begin()
	t.Logf("%+v\n", ao)
}
