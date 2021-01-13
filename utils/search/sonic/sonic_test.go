package sonic

import (
	"reflect"
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
		{Object: "id:2m2n3b4vf6", Text: "好大一条狗"},
		{Object: "id:3m2n3b4vf6", Text: "好大一只猫"},
	}); err != nil {
		t.Fatal(err)
	}

	for _, term := range []struct {
		term string
		want []string
	}{
		{term: "吃饭", want: []string{"id:6ab56b4kk3"}},
		{term: "好人", want: []string{"id:5hg67f8dg5"}},
		{term: "天气", want: []string{"id:1m2n3b4vf6"}},
		{term: "哈哈", want: []string{""}}, // 没有内容也要返回一个空字符串元素，是为什么呢？
		{term: "天天", want: []string{"id:1m2n3b4vf6"}},

		//
		{term: "昨气", want: []string{"id:1m2n3b4vf6"}},
		{term: "昨好", want: []string{"id:1m2n3b4vf6"}},

		{term: "好大", want: []string{"id:2m2n3b4vf6", "id:3m2n3b4vf6"}},
	} {
		r, err := client.Query(collec, bucket, term.term, 10, 0)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("term: %s, r: %+v, len: %v\n", term.term, r, len(r))
		if len(r) != 0 && !reflect.DeepEqual(r, term.want) {
			t.Fatalf("Bad result: %v != %v\n", r, term.want)
		}
	}
}
