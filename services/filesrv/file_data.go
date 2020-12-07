package filesrv

import (
	"github.com/donnol/jdnote/models/commonmodel"
	"github.com/donnol/jdnote/models/filemodel"
)

type GetParam struct {
	commonmodel.IDBase
}

type GetResult struct {
	filemodel.FileData
	filemodel.FileContentData
}

type AddParam struct {
	filemodel.FileInputData
	filemodel.FileContentData
}

type AddResult struct {
	commonmodel.IDBase

	Path string `json:"path"` // 路径，如：/v1/file
}
