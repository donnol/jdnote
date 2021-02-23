package filestore

import (
	"context"

	"github.com/donnol/jdnote/models/filemodel"
)

type IFile interface {
	Add(ctx context.Context, entity filemodel.File) (id int, err error)
	AddContent(ctx context.Context, entity filemodel.FileContent) (id int, err error)
	Get(ctx context.Context, id int) (entity filemodel.File, err error)
	GetContentByIDs(ctx context.Context, ids []int64) (entity []filemodel.FileContent, err error)
}
