package filesrv

import (
	"github.com/donnol/jdnote/models/filemodel"
	"github.com/donnol/jdnote/stores/filestore"
	"github.com/donnol/jdnote/utils/context"
	"github.com/pkg/errors"
)

func NewIFile(
	fileStore filestore.IFile,
) IFile {
	return &fileImpl{
		fileStore: fileStore,
	}
}

type fileImpl struct {
	fileStore filestore.IFile
}

func (impl *fileImpl) Get(ctx context.Context, param GetParam) (result GetResult, err error) {
	file, err := impl.fileStore.Get(ctx, param.ID)
	if err != nil {
		return
	}
	contents, err := impl.fileStore.GetContentByIDs(ctx, []int64{int64(file.FileContentID)})
	if err != nil {
		return
	}
	if len(contents) == 0 {
		return result, errors.Errorf("找不到文件内容")
	}
	result.FileContentData = contents[0].FileContentData
	result.FileData = file.FileData

	return
}

func (impl *fileImpl) Add(ctx context.Context, param AddParam) (result AddResult, err error) {
	if err = context.WithTx(ctx, func(ctx context.Context) error {
		var fileContentID int
		fileContentID, err = impl.fileStore.AddContent(ctx, filemodel.FileContent{FileContentData: filemodel.FileContentData{Content: param.Content}})
		if err != nil {
			return err
		}

		var size int64 = int64(len(param.Content))
		result.ID, err = impl.fileStore.Add(ctx, filemodel.File{
			FileRelation: filemodel.FileRelation{
				FileContentID: fileContentID,
			},
			FileData: filemodel.FileData{
				FileInputData: param.FileInputData,
				FileDeriveData: filemodel.FileDeriveData{
					Size: size,
				},
			},
		})
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return
	}
	return
}
