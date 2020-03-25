package models

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/donnol/jdnote/utils/context"
	utillog "github.com/donnol/jdnote/utils/log"
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
	// ctx.Logger().Debugf("scan: %+v\n", row)

	ei.rows = append(ei.rows, row)

	return nil
}

// Do 处理
func (ei *entityImpl) Do(ctx context.Context) error {
	// 处理数据，这里需要用到ctx
	// 比如在标题后添加后缀
	// suffix := "_hah"

	// 这里换成批量操作，会更快
	for _, single := range ei.rows {
		// title := single.Title+suffix
		title := "title4"
		res := ctx.DB().MustExecContext(ctx, "UPDATE t_note SET title = $1 WHERE id = $2", title, single.ID)
		afNum, err := res.RowsAffected()
		if err != nil {
			return err
		}
		_ = afNum
		// ctx.Logger().Debugf("affected row is %d, id is %v\n", afNum, single.ID)
	}

	return nil
}

func TestBaseProcess(t *testing.T) {
	ctx := context.New(defaultDB, utillog.New(os.Stdout, "", log.LstdFlags), 0)
	base := NewBase()

	entity := &entityImpl{}
	opt := ProcessOption{
		Query: "SELECT id, title FROM t_note",
		Args:  []interface{}{},

		N: 10, // 批数量 2(60s) 100(25s) 200(25s) 300(23s) 400(25s) 500(24s)
		Entity: func() Entity {
			return &entityImpl{
				rows: make([]rowStruct, 0),
			}
		},
	}
	if err := base.ProcessConcurrent(ctx, opt); err != nil {
		t.Fatal(err)
	}

	t.Logf("entity: %+v\n", entity)
}
