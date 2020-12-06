package fileapi

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"

	"github.com/donnol/jdnote/app"
)

func setRequestCookie(r *http.Request) error {
	ctx := context.Background()
	appObj, _ := app.New(ctx)
	cookie, err := appObj.MakeCookie(1)
	if err != nil {
		return (err)
	}
	r.AddCookie(&cookie)

	return nil
}

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

	fileName2 := "testdata/testcsv.csv"
	filePath2 := path.Join(fileDir, fileName2)

	// 打开文件
	file2, err := os.Open(filePath2)
	if err != nil {
		t.Fatal(err)
	}
	defer file2.Close()

	// 新建multipart writer，将文件内容赋值到body里
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 第一个文件
	part, err := writer.CreateFormFile("file_md", filepath.Base(file.Name()))
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatal(err)
	}

	// 第二个文件
	part2, err := writer.CreateFormFile("file_csv", filepath.Base(file2.Name()))
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(part2, file2)
	if err != nil {
		t.Fatal(err)
	}

	// 写入参数
	err = writer.WriteField("fieldFileName", filepath.Base(file.Name()))
	if err != nil {
		t.Fatal(err)
	}
	writer.Close()

	t.Logf("%s\n", body.Bytes())

	// 新建请求
	r, err := http.NewRequest("POST", "http://127.0.0.1:8890/v1/file", body)
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Add("Content-Type", writer.FormDataContentType())
	setRequestCookie(r)

	// 发送请求
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", data)

	// 结果打印
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Bad status code: %d\n", resp.StatusCode)
	}
}

func TestGetFile(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8890/v1/file?id=2", nil)
	if err != nil {
		t.Fatal(err)
	}
	// 不加这个hader，会报"unexpected EOF"错误
	// 了解 https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Accept-Encoding
	r.Header.Add("Accept-Encoding", "identity")
	if err := setRequestCookie(r); err != nil {
		t.Fatal(err)
	}

	// 发送请求
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", data)
	// 结果打印
	if resp.StatusCode != http.StatusOK {
		t.Fatal("Bad status code")
	}
}
