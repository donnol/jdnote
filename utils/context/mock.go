package context

import (
	"time"

	"github.com/donnol/jdnote/utils/store/db"
)

type ContextMock struct {
	DBFunc func() db.DB

	DeadlineFunc func() (deadline time.Time, ok bool)

	DoneFunc func() <-chan struct{}

	ErrFunc func() error

	NewWithTxFunc func(db.DB) Context

	ValueFunc func(key interface{}) interface{}
}

var _ Context = &ContextMock{}

func (mockRecv *ContextMock) DB() db.DB {
	return mockRecv.DBFunc()
}

func (mockRecv *ContextMock) Deadline() (deadline time.Time, ok bool) {
	return mockRecv.DeadlineFunc()
}

func (mockRecv *ContextMock) Done() <-chan struct{} {
	return mockRecv.DoneFunc()
}

func (mockRecv *ContextMock) Err() error {
	return mockRecv.ErrFunc()
}

func (mockRecv *ContextMock) NewWithTx(p0 db.DB) Context {
	return mockRecv.NewWithTxFunc(p0)
}

func (mockRecv *ContextMock) Value(key interface{}) interface{} {
	return mockRecv.ValueFunc(key)
}
