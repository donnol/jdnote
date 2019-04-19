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
	fileDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	fileName := "testdata/testfile.md"
	filePath := path.Join(fileDir, fileName)

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("fileName", filepath.Base(file.Name()))
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatal(err)
	}
	err = writer.WriteField("fileName", filepath.Base(file.Name()))
	if err != nil {
		t.Fatal(err)
	}
	defer writer.Close()

	t.Logf("%s\n", body.Bytes())
	r, err := http.NewRequest("POST", "http://127.0.0.1:8810/v1/file", body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v\n", writer.FormDataContentType())
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	t.Log(resp.StatusCode)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", data)
}
