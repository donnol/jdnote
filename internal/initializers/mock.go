package initializers

import (
	"database/sql"

	"github.com/donnol/jdnote/utils/context"
)

type EntityMock struct {
	DoFunc func(context.Context) error

	ScanFunc func(context.Context, *sql.Rows) error
}

var _ Entity = &EntityMock{}

func (mockRecv *EntityMock) Do(p0 context.Context) error {
	return mockRecv.DoFunc(p0)
}

func (mockRecv *EntityMock) Scan(p0 context.Context, p1 *sql.Rows) error {
	return mockRecv.ScanFunc(p0, p1)
}

type ScannerMock struct {
	ScanFunc func(context.Context, *sql.Rows) error
}

var _ Scanner = &ScannerMock{}

func (mockRecv *ScannerMock) Scan(p0 context.Context, p1 *sql.Rows) error {
	return mockRecv.ScanFunc(p0, p1)
}

type DoerMock struct {
	DoFunc func(context.Context) error
}

var _ Doer = &DoerMock{}

func (mockRecv *DoerMock) Do(p0 context.Context) error {
	return mockRecv.DoFunc(p0)
}
