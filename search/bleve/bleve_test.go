package bleve_test

import (
	"testing"

	"github.com/blevesearch/bleve"
)

func TestBleve(t *testing.T) {
	// 构建
	message := struct {
		Id   string
		From string
		Body string
	}{
		Id:   "example",
		From: "marty.schoch@gmail.com",
		Body: "bleve indexing is easy",
	}

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("jdscript.com", mapping)
	if err != nil {
		panic(err)
	}
	index.Index(message.Id, message)

	// 查询
	nindex, err := bleve.Open("jdscript.com")
	if err != nil {
		t.Fatal(err)
	}
	query := bleve.NewQueryStringQuery("bleve")
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := nindex.Search(searchRequest)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", searchResult)
}
