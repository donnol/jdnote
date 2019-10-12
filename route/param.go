package route

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/donnol/jdnote/utils/context"
	"github.com/pkg/errors"
)

// Param 参数
type Param struct {
	// 方法
	method string

	// 参数
	body   []byte
	values url.Values

	// 文件
	multipartReader *multipart.Reader
}

// Parse 解析
func (p *Param) Parse(ctx context.Context, v interface{}) error {
	var err error

	// 解析
	switch p.method {
	case http.MethodPost:
		fallthrough
	case http.MethodPut:
		err = json.Unmarshal(p.body, v)
	case http.MethodGet:
		fallthrough
	case http.MethodDelete:
		err = decoder.Decode(v, p.values)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	// 检查参数
	if vv, ok := v.(Checker); ok {
		if err := vv.Check(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// ParseMultipartForm 解析内容
func (p *Param) ParseMultipartForm(maxFileSize int64, v interface{}) ([]byte, error) {
	var body []byte

	if p.multipartReader == nil {
		return body, fmt.Errorf("Bad multipart reader")
	}

	// 使用ReadForm
	form, err := p.multipartReader.ReadForm(maxFileSize)
	if err != nil {
		return body, err
	}

	// 获取参数
	if err := decoder.Decode(v, form.Value); err != nil {
		return body, err
	}

	// 获取内容
	buf := new(bytes.Buffer)
	for _, single := range form.File {
		for _, one := range single {
			file, err := one.Open()
			if err != nil {
				return body, err
			}
			defer file.Close()

			_, err = buf.ReadFrom(file)
			if err != nil {
				return body, err
			}
		}
	}
	body = buf.Bytes()

	return body, nil
}
