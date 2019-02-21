package jwt

import (
	"math/rand"
	"testing"
)

func TestToken(t *testing.T) {
	token := &Token{}
	id := rand.Int()
	r, err := token.Sign(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)

	userID, err := token.Verify(r)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(userID)

	if userID != id {
		t.Fatalf("Bad userID, got %d\n", userID)
	}
}
