package process

import (
	stdctx "context"
	"database/sql"
	"math/rand"
	"strconv"
	"testing"

	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/store/db"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

type rowStruct struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}

type entityImpl struct {
	rows []rowStruct
}

// Scan 扫描
func (ei *entityImpl) Scan(ctx context.Context, rows *sql.Rows) error {
	var row rowStruct

	// 将数据写到具体的结构体中
	if err := rows.Scan(&row.ID, &row.Title); err != nil {
		return err
	}

	ei.rows = append(ei.rows, row)

	return nil
}

// Do 处理
func (ei *entityImpl) Do(ctx context.Context) error {
	ids := make([]int, 0, len(ei.rows))
	for _, single := range ei.rows {
		ids = append(ids, single.ID)
	}

	title := "title" + strconv.Itoa(rand.Int())
	query, args, err := sqlx.In("UPDATE t_note SET title = ? WHERE id IN (?)", title, ids)
	if err != nil {
		return errors.WithStack(err)
	}
	query = ctx.DB().Rebind(query)
	res := ctx.DB().MustExecContext(ctx, query, args...)
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

var (
	sctx = stdctx.Background()
	ctx  = func() context.Context {
		dbImpl, err := db.Open(db.Option{
			DriverName:     "postgres",
			DataSourceName: "postgres://jd:jd@127.0.0.1:5432/cicada_dev?sslmode=disable",
		})
		if err != nil {
			panic(err)
		}
		return context.New(sctx, dbImpl)
	}()
	batchProcessor = NewBatchProcessor()

	entity = &entityImpl{
		rows: make([]rowStruct, 0),
	}
	opt = ProcessOption{
		Query: "SELECT id, title FROM t_note",
		Args:  []interface{}{},

		N: 10, // 批数量 2(60s) 100(25s) 200(25s) 300(23s) 400(25s) 500(24s)
		NewEntity: func() Entity {
			return entity
		},
	}
)

func TestBaseProcess(t *testing.T) {
	if err := batchProcessor.ProcessConcurrent(ctx, opt); err != nil {
		t.Fatal(err)
	}

	t.Logf("entity: %+v\n", entity)
}

func BenchmarkBaseProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = batchProcessor.ProcessConcurrent(ctx, opt)
	}
}
