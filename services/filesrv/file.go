package filesrv

import (
	"github.com/donnol/jdnote/stores/filestore"
	"github.com/donnol/jdnote/utils/context"
)

func NewIFile() IFile {
	return &fileImpl{}
}

type fileImpl struct {
	fileStore filestore.IFile
}

func (impl *fileImpl) Get(ctx context.Context, param GetParam) (result GetResult, err error) {
	return
}

func (impl *fileImpl) Add(ctx context.Context, param AddParam) (result AddResult, err error) {
	return
}
