package generate

import (
	"bytes"
	"go/build"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

// MakeTemplate 生成模板
func MakeTemplate(temp string, data map[string]interface{}) []byte {
	temp = `// Please don't edit this file, it's made by go generate!
	` + temp

	// 解析
	t, err := template.New("").Parse(temp)
	if err != nil {
		log.Fatal(err)
	}

	// 写入到buff
	buff := new(bytes.Buffer)
	err = t.Execute(buff, data)
	if err != nil {
		log.Fatal(err)
	}

	return buff.Bytes()
}

// WrapTemplate 封装模板
func WrapTemplate(f func(pkgName string) []byte) []byte {
	// 获取包名
	pkgInfo, err := build.ImportDir(".", 0)
	if err != nil {
		log.Fatal(err)
	}

	// 获取模板
	content := f(pkgInfo.Name)

	// 格式化
	formatContent, err := format.Source([]byte(content))
	if err != nil {
		log.Fatal(err)
	}

	return formatContent
}

// Save 保存
func Save(fileName string, content []byte) {
	if err := ioutil.WriteFile(fileName, content, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
