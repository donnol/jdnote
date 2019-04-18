package file

import (
	"github.com/donnol/jdnote/api"
	"github.com/donnol/jdnote/route"
)

func init() {
	route.DefaultRouter.Register(&File{})
}

// File 文件
type File struct {
	api.Base

	// 所属的Group
	V1 route.Group // 属于v1分组
}

// Add 上传文件
func (file *File) Add(param route.Param) (r route.Result, err error) {
	// TODO: 这里要告诉route我是一个文件上传接口
	// 直接用File？还是在File里添加指定类型字段呢？这样做是针对整个结构体来说的
	// 如果只针对单个方法呢？在方法名后加后缀？类似：Add_File？

	return
}

// Get 下载文件
func (file *File) Get(param route.Param) (r route.Result, err error) {

	return
}
