package file

import (
	"bytes"

	"github.com/donnol/jdnote/api"
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/route"
)

func init() {
	route.Register(&File{})
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
func (file *File) Add(ctx context.Context, param route.Param) (r route.Result) {
	p := struct {
		FileName string `json:"fileName"`
	}{}
	body, err := param.ParseMultipartForm(64*1024*1024, &p)
	r.SetErr(err)
	ctx.Logger().Debugf("%+v, %d\n", p, len(body))

	return
}

// Get 下载文件
func (file *File) Get(ctx context.Context, param route.Param) (r route.Result) {
	// 参数
	r.SetErr(param.Parse(&struct{}{}))

	// 权限
	_ = ctx.UserID()

	// 业务
	filename := "test.md"
	content := "# Hello\n\n## I am bat man\n\n"

	buf := new(bytes.Buffer)
	_, err := buf.Write([]byte(content))
	r.SetErr(err)
	r.Content = route.MakeContentFromBuffer(filename, buf)
	ctx.Logger().Debugf("r: %+v\n", r)

	return
}
