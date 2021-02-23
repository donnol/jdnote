package filestore

import (
	"context"

	"github.com/donnol/jdnote/models/filemodel"
	"github.com/donnol/jdnote/utils/store/db"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

func NewIFile(
	db db.DB,
) IFile {
	return &fileImpl{
		db: db,
	}
}

type fileImpl struct {
	db db.DB
}

func (impl *fileImpl) Get(ctx context.Context, id int) (entity filemodel.File, err error) {
	err = db.DBFromCtxValue(ctx, impl.db).GetContext(ctx, &entity, `
		SELECT *
		FROM t_file
		WHERE id = $1
		`,
		id,
	)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (impl *fileImpl) Add(ctx context.Context, entity filemodel.File) (id int, err error) {
	err = db.DBFromCtxValue(ctx, impl.db).GetContext(ctx, &id, `INSERT INTO t_file(file_content_id, name, size)
		VALUES($1, $2, $3)
		RETURNING id
		`,
		entity.FileContentID,
		entity.Name,
		entity.Size,
	)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (impl *fileImpl) GetContentByIDs(ctx context.Context, ids []int64) (entity []filemodel.FileContent, err error) {
	err = db.DBFromCtxValue(ctx, impl.db).SelectContext(ctx, &entity, `
		SELECT *
		FROM t_file_content
		WHERE id = any($1)
		ORDER BY id DESC
		`,
		pq.Int64Array(ids),
	)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (impl *fileImpl) AddContent(ctx context.Context, entity filemodel.FileContent) (id int, err error) {
	err = db.DBFromCtxValue(ctx, impl.db).GetContext(ctx, &id, `INSERT INTO t_file_content(content)
		VALUES($1)
		RETURNING id
		`,
		entity.Content,
	)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
