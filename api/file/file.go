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

	return
}

// Get 下载文件
func (file *File) Get(param route.Param) (r route.Result, err error) {

	return
}
