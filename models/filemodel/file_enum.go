package filemodel

// Ext 文件扩展名: doc 文档;png png图片;ppt ppt文件;
type Ext string

const (
	ExtDoc Ext = "doc"
	ExtPng Ext = "png"
	ExtPpt Ext = "ppt"
)

var (
	allExt = [...]Ext{
		ExtDoc,
		ExtPng,
		ExtPpt,
	}
)

func (ext Ext) IsValid() bool {
	for _, one := range allExt {
		if one == ext {
			return true
		}
	}
	return false
}
