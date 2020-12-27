package sonic

import (
	"testing"

	"github.com/expectedsh/go-sonic/sonic"
)

func TestSonic(t *testing.T) {
	host := "localhost"
	port := 1491

	client, err := New(Option{
		Host:     host,
		Port:     port,
		Password: "jdis1gHR",
	})
	if err != nil {
		t.Fatal(err)
	}

	collec := "mycollec"
	bucket := "mybucket"

	if err := client.BulkPush(collec, bucket, 3, []sonic.IngestBulkRecord{
		{Object: "id:6ab56b4kk3", Text: "星期一晚上去吃饭"},
		{Object: "id:5hg67f8dg5", Text: "我是一个好人"},
		{Object: "id:1m2n3b4vf6", Text: "昨天天气很好"},
	}); err != nil {
		t.Fatal(err)
	}

	for _, term := range []string{
		"吃饭",
		"好人",
		"天气",
		"哈哈",
	} {
		r, err := client.Query(collec, bucket, term, 10, 0)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("term: %s, r: %+v\n", term, r)
	}
}
