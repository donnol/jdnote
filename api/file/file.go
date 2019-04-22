package file

import (
	"bytes"
	"mime/multipart"

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

	// 标志文件上传/下载属性
	File route.File `file:"Add,get"` // 使用tag来指定方法
}

// Add 上传文件
func (file *File) Add(param route.Param) (r route.Result, err error) {
	p := struct {
		FileName string `json:"fileName"`
	}{}
	body, err := param.ParseMultipartForm(64*1024*1024, &p)
	if err != nil {
		return
	}
	file.Debugf("%+v, %d\n", p, len(body))

	return
}

// Get 下载文件
func (file *File) Get(param route.Param) (r route.Result, err error) {
	// 参数
	if err = param.Parse(&struct{}{}); err != nil {
		return
	}

	// 权限
	_ = param.UserID

	// 业务
	content := "# Hello\n\n## I am bat man\n\n"
	buf := new(bytes.Buffer)
	_, err = buf.Write([]byte(content))
	if err != nil {
		return
	}
	writer := multipart.NewWriter(buf)
	r.ContentLength = int64(buf.Len())
	r.ContentType = writer.FormDataContentType()
	r.ContentReader = buf
	r.ExtraHeaders = map[string]string{
		"Content-Disposition": `attachment; filename="test.md"`,
	}
	file.Debugf("r: %+v\n", r)

	return
}
