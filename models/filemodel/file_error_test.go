package filemodel_test

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/donnol/jdnote/models/filemodel"
)

func TestIsErrNil(t *testing.T) {
	for _, cas := range []struct {
		err  error
		want bool
	}{
		{err: filemodel.ErrNotFound, want: false},
		{err: nil, want: true},
		{err: io.EOF, want: false},
	} {
		r := cas.err == nil
		if r != cas.want {
			t.Fatalf("Bad result, %v != %v\n", r, cas.want)
		}
	}
}

func TestIsErrNotFound(t *testing.T) {
	for _, cas := range []struct {
		err  error
		want bool
	}{
		{err: filemodel.ErrNotFound, want: true},
		{err: nil, want: false},
		{err: errors.New("abc"), want: false},
		{err: fmt.Errorf("abc"), want: false},
		{err: io.EOF, want: false},
	} {
		r := filemodel.IsErrNotFound(cas.err)
		if r != cas.want {
			t.Fatalf("Bad result, %v != %v\n", r, cas.want)
		}
	}
}
