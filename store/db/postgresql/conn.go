package pg

import (
	"github.com/donnol/jdnote/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // github.com/lib/pq postgresql驱动
)

// defaultDB 默认db
var defaultDB = func() *sqlx.DB {

	db, err := sqlx.Open(config.Default().DB.Scheme, config.Default().DB.String())
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}()
