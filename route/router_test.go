package route

import (
	"testing"
)

func getUser(p Param) (r Result, err error) {

	return
}

func getUserCurrent(p Param) (r Result, err error) {

	return
}

func TestGetMethodPathFromFunc(t *testing.T) {
	method, path := getMethodPathFromFunc(getUser)
	t.Log(method, path)

	method, path = getMethodPathFromFunc(getUserCurrent)
	t.Log(method, path)
}
