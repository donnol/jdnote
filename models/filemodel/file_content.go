package filemodel

import "github.com/donnol/jdnote/models/commonmodel"

// FileContent 文件内容
type FileContent struct {
	commonmodel.Base

	FileContentData
}

type FileContentData struct {
	Content []byte `json:"content"` // 内容
}
