package action

import (
	"testing"

	pg "github.com/donnol/jdnote/store/db/postgresql"
)

func TestGet(t *testing.T) {
	a := &Action{
		DB:     pg.New(),
		Action: "ALL",
	}
	if err := a.Add(); err != nil {
		t.Fatal(err)
	}

	if err := a.Get(1); err != nil {
		t.Fatal(err)
	}
	t.Log(a)
}
