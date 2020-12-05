package filesrv

import "github.com/donnol/jdnote/models/filemodel"

type GetParam struct {
	ID int `json:"id"` // 记录ID
}

type GetResult struct {
	filemodel.FileData
	filemodel.FileContentData
}

type AddParam struct {
}

type AddResult struct {
}
