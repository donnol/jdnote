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

func TestBaseProcess(t *testing.T) {
	ctx := context.New(defaultDB, utillog.New(os.Stdout, "", log.LstdFlags), 0)
	base := NewBase()
	opt := ProcessOption{
		Query: "SELECT id, title FROM t_note ORDER BY id ASC",
		Args:  []interface{}{},

		N: 2, // 批数量
		Scan: func(ctx context.Context, rows *sql.Rows) (interface{}, error) {
			var row rowStruct

			// 将数据写到具体的结构体中
			if err := rows.Scan(&row.ID, &row.Title); err != nil {
				return row, err
			}
			ctx.Logger().Debugf("scan: %+v\n", row)

			return row, nil
		}, // 扫描
		Do: func(ctx context.Context, data []interface{}) error {
			// 这里要遍历一次，逐个断言才能拿到具体信息
			var rows = make([]rowStruct, 0, len(data))
			for _, single := range data {
				tmp := single.(rowStruct)
				ctx.Logger().Debugf("data is %+v\n", tmp)

				rows = append(rows, tmp)
			}

			// 处理数据，这里需要用到ctx
			// 比如在标题后添加后缀
			suffix := "_hah"
			for _, single := range rows {
				res := ctx.DB().MustExecContext(ctx, "UPDATE t_note SET title = $1 WHERE id = $2", single.Title+suffix, single.ID)
				afNum, err := res.RowsAffected()
				if err != nil {
					return err
				}
				ctx.Logger().Debugf("affected row is %d\n", afNum)
			}

			return nil
		}, // 处理
	}
	if err := base.ProcessConcurrent(ctx, opt); err != nil {
		t.Fatal(err)
	}
}
