package log

import (
	"errors"
	"testing"
)

func TestInfo(t *testing.T) {
	for _, cas := range []struct {
		value interface{}
	}{
		{errors.New("Test Message")},
	} {
		Fatalf("%+v\n", cas.value)
		Errorf("%+v\n", cas.value)
		Warnf("%+v\n", cas.value)
		Infof("%+v\n", cas.value)
		Debugf("%+v\n", cas.value)
		Tracef("%+v\n", cas.value)
	}
}
