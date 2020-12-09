package filestore

import (
	"github.com/donnol/jdnote/models/filemodel"
	"github.com/donnol/jdnote/utils/context"
)

type FileMock struct {
	AddFunc func(ctx context.Context, entity filemodel.File) (id int, err error)

	AddContentFunc func(ctx context.Context, entity filemodel.FileContent) (id int, err error)

	GetFunc func(ctx context.Context, id int) (entity filemodel.File, err error)

	GetContentByIDsFunc func(ctx context.Context, ids []int64) (entity []filemodel.FileContent, err error)
}

var _ IFile = &FileMock{}

func (mockRecv *FileMock) Add(ctx context.Context, entity filemodel.File) (id int, err error) {
	return mockRecv.AddFunc(ctx, entity)
}

func (mockRecv *FileMock) AddContent(ctx context.Context, entity filemodel.FileContent) (id int, err error) {
	return mockRecv.AddContentFunc(ctx, entity)
}

func (mockRecv *FileMock) Get(ctx context.Context, id int) (entity filemodel.File, err error) {
	return mockRecv.GetFunc(ctx, id)
}

func (mockRecv *FileMock) GetContentByIDs(ctx context.Context, ids []int64) (entity []filemodel.FileContent, err error) {
	return mockRecv.GetContentByIDsFunc(ctx, ids)
}
