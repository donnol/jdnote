package filemodel

import "github.com/donnol/jdnote/models/commonmodel"

// File 文件，对应表，包括公有字段和特有字段
// 适合在store层使用，因为一般整表使用
// 一个表一般来说，会包括几个部分，分别有：公用字段、关联字段、数据字段（又可以分为输入数据和衍生数据）
type File struct {
	commonmodel.TableBase

	FileRelation
	FileData
}

// FileRelation 关联字段，对应表外键
// 有些表有，有些表没有
type FileRelation struct {
	FileContentID int `json:"fileContentID" db:"file_content_id"` // 内容ID
}

// FileData 特有数据，对应表特有字段
// 适合在service层使用，因为一般需要重新组装
// 也因为用了Data后缀，所以表名请不要使用xxx_data这种了
type FileData struct {
	FileInputData
	FileDeriveData
}

type FileInputData struct {
	Name string `json:"name"` // 名称
}

type FileDeriveData struct {
	Size int64 `json:"size"` // 大小
}

func (fd FileData) Filter() interface{} {

	return fd
}
