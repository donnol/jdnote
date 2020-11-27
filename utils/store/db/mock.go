package db

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DBMock struct {
	BindNamedFunc func(query string, arg interface{}) (string, []interface{}, error)

	DriverNameFunc func() string

	GetFunc func(interface{}, string, ...interface{}) error

	GetContextFunc func(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	MustExecFunc func(query string, args ...interface{}) sql.Result

	MustExecContextFunc func(ctx context.Context, query string, args ...interface{}) sql.Result

	NamedExecFunc func(query string, arg interface{}) (sql.Result, error)

	NamedExecContextFunc func(ctx context.Context, query string, arg interface{}) (sql.Result, error)

	NamedQueryFunc func(query string, arg interface{}) (*sqlx.Rows, error)

	PrepareNamedFunc func(query string) (*sqlx.NamedStmt, error)

	PrepareNamedContextFunc func(ctx context.Context, query string) (*sqlx.NamedStmt, error)

	PreparexFunc func(query string) (*sqlx.Stmt, error)

	PreparexContextFunc func(ctx context.Context, query string) (*sqlx.Stmt, error)

	QueryContextFunc func(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

	QueryRowxFunc func(query string, args ...interface{}) *sqlx.Row

	QueryRowxContextFunc func(ctx context.Context, query string, args ...interface{}) *sqlx.Row

	QueryxFunc func(query string, args ...interface{}) (*sqlx.Rows, error)

	QueryxContextFunc func(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)

	RebindFunc func(query string) string

	SelectFunc func(dest interface{}, query string, args ...interface{}) error

	SelectContextFunc func(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

var _ DB = &DBMock{}

func (mockRecv *DBMock) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	return mockRecv.BindNamedFunc(query, arg)
}

func (mockRecv *DBMock) DriverName() string {
	return mockRecv.DriverNameFunc()
}

func (mockRecv *DBMock) Get(p0 interface{}, p1 string, p2 ...interface{}) error {
	return mockRecv.GetFunc(p0, p1, p2...)
}

func (mockRecv *DBMock) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return mockRecv.GetContextFunc(ctx, dest, query, args...)
}

func (mockRecv *DBMock) MustExec(query string, args ...interface{}) sql.Result {
	return mockRecv.MustExecFunc(query, args...)
}

func (mockRecv *DBMock) MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result {
	return mockRecv.MustExecContextFunc(ctx, query, args...)
}

func (mockRecv *DBMock) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return mockRecv.NamedExecFunc(query, arg)
}

func (mockRecv *DBMock) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return mockRecv.NamedExecContextFunc(ctx, query, arg)
}

func (mockRecv *DBMock) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	return mockRecv.NamedQueryFunc(query, arg)
}

func (mockRecv *DBMock) PrepareNamed(query string) (*sqlx.NamedStmt, error) {
	return mockRecv.PrepareNamedFunc(query)
}

func (mockRecv *DBMock) PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error) {
	return mockRecv.PrepareNamedContextFunc(ctx, query)
}

func (mockRecv *DBMock) Preparex(query string) (*sqlx.Stmt, error) {
	return mockRecv.PreparexFunc(query)
}

func (mockRecv *DBMock) PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error) {
	return mockRecv.PreparexContextFunc(ctx, query)
}

func (mockRecv *DBMock) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return mockRecv.QueryContextFunc(ctx, query, args...)
}

func (mockRecv *DBMock) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return mockRecv.QueryRowxFunc(query, args...)
}

func (mockRecv *DBMock) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return mockRecv.QueryRowxContextFunc(ctx, query, args...)
}

func (mockRecv *DBMock) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return mockRecv.QueryxFunc(query, args...)
}

func (mockRecv *DBMock) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return mockRecv.QueryxContextFunc(ctx, query, args...)
}

func (mockRecv *DBMock) Rebind(query string) string {
	return mockRecv.RebindFunc(query)
}

func (mockRecv *DBMock) Select(dest interface{}, query string, args ...interface{}) error {
	return mockRecv.SelectFunc(dest, query, args...)
}

func (mockRecv *DBMock) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return mockRecv.SelectContextFunc(ctx, dest, query, args...)
}
