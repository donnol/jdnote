package minio

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	// 请先开启一个测试服务
	opt := Option{
		Endpoint:        "127.0.0.1:9000",
		AccessKeyID:     "minioadmin",
		SecretAccessKey: "minioadmin",
		UseSSL:          false,
	}
	client, err := New(opt)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	obj, err := client.GetObject(ctx, GetOption{
		BucketName: "first",
		ObjectName: "one/hangye-yinyue.jpg",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("obj: %+v\n", obj)

	buf := new(bytes.Buffer)
	n, err := buf.ReadFrom(obj)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("n: %d, buf: %v\n", n, buf.Len())

	var fileReader io.Reader
	fileReader, err = os.OpenFile("testdata/test.jpg", os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	objSize := int64(buf.Len())
	contentType := "image/jpeg"
	r, err := client.PutObject(ctx, PutOption{
		BucketName:  "first",
		ObjectName:  "one/test.jpg",
		Reader:      fileReader,
		ObjectSize:  objSize,
		ContentType: contentType,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("r: %+v\n", r)
}
