package fileapi

import (
	"bytes"

	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/utils/route"
	"github.com/donnol/tools/log"
)

// File 文件
type File struct {
	// 所属的Group
	V1 route.Group // 属于v1分组

	// 标志文件上传/下载属性
	File route.File `file:"Add,get"` // 使用tag来指定方法

	// 有tag则只对tag里的添加，没有则全部方法均添加(这时参数怎么指定呢？`rate:"Rate(0.25, 2)"`)
	Limiter route.Limiter `method:"Add(0.25, 2);Get(0.25, 2)"` // 指定限流器，包括方法和参数; 多个方法使用分号分隔

	logger log.Logger
}

// Add 上传文件
func (file *File) Add(ctx context.Context, param route.Param) (r route.Result, err error) {
	p := struct {
		FieldFileName string `json:"fieldFileName"`
	}{}
	body, err := param.ParseMultipartForm(64*1024*1024, &p)
	if err != nil {
		return
	}
	file.logger.Debugf("%+v, %d\n", p, len(body))
	for name, one := range body {
		file.logger.Debugf("name: %s, content: %s\n", name, one)
	}

	return
}

// Get 下载文件
func (file *File) Get(ctx context.Context, param route.Param) (r route.Result, err error) {
	// 参数
	if err = param.Parse(ctx, &struct{}{}); err != nil {
		return
	}

	// 权限

	// 业务
	filename := "test.md"
	content := "# Hello\n\n## I am bat man\n\n"

	buf := new(bytes.Buffer)
	_, err = buf.Write([]byte(content))
	if err != nil {
		return
	}
	r.Content = route.MakeContentFromBuffer(filename, buf)
	file.logger.Debugf("r: %+v\n", r)

	return
}
