package sonic

import (
	"fmt"

	"github.com/expectedsh/go-sonic/sonic"
)

type Client struct {
	ingester sonic.Ingestable
	search   sonic.Searchable
}

type Option struct {
	Host     string
	Port     int
	Password string
}

func New(opt Option) (*Client, error) {
	client := &Client{}

	var err error
	client.ingester, err = sonic.NewIngester(opt.Host, opt.Port, opt.Password)
	if err != nil {
		return nil, err
	}

	client.search, err = sonic.NewSearch(opt.Host, opt.Port, opt.Password)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (client *Client) BulkPush(collection string, bucket string, parallelRoutines int, records []sonic.IngestBulkRecord) error {

	var err error
	pushResults := client.ingester.BulkPush(collection, bucket, parallelRoutines, records)
	for i, one := range pushResults {
		if one.Error != nil {
			if err != nil {
				err = fmt.Errorf("No.%d: %+v, err is %+v, old is %+v", i, one, one.Error, err)
			} else {
				err = fmt.Errorf("No.%d: %+v, err is %+v", i, one, one.Error)
			}
		}
	}
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) Query(collection string, bucket string, terms string, limit int, offset int) (results []string, err error) {
	results, err = client.search.Query(collection, bucket, terms, limit, offset)
	if err != nil {
		return
	}

	return
}
