package file

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"
)

func TestAddFile(t *testing.T) {
	// 获取运行目录
	fileDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	// 获取文件名和文件路径
	fileName := "testdata/testfile.md"
	filePath := path.Join(fileDir, fileName)

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// 新建multipart writer，将文件内容赋值到body里
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatal(err)
	}
	// 写入参数
	err = writer.WriteField("fileName", filepath.Base(file.Name()))
	if err != nil {
		t.Fatal(err)
	}
	writer.Close()

	// 新建请求
	t.Logf("%s\n", body.Bytes())
	r, err := http.NewRequest("POST", "http://127.0.0.1:8810/v1/file", body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v\n", writer.Boundary())
	r.Header.Add("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// 结果打印
	t.Log(resp.StatusCode)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", data)
}
