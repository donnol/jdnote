package filemodel

import "github.com/donnol/jdnote/models/commonmodel"

// File 文件，对应表，包括公有字段和特有字段
// 适合在store层使用，因为一般整表使用
type File struct {
	commonmodel.Base

	FileData
}

// FileData 特有数据，对应表特有字段
// 适合在service层使用，因为一般需要重新组装
// 也因为用了Data后缀，所以表名请不要使用xxx_data这种了
type FileData struct {
	FileContentID int    `json:"fileContentID" db:"file_content_id"` // 内容ID
	Name          string `json:"name"`                               // 名称
	Size          int64  `json:"size"`                               // 大小
}
