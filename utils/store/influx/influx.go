package influx

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type Client struct {
	influxdb2.Client
}

type Option struct {
	Host  string // 如：http://localhost:8086
	Token string // 如：xxxx-abc-efg
}

type BucketSetting struct {
	OrgName    string
	BucketName string
}

func Open(opt Option, infOpts *influxdb2.Options) *Client {
	var client influxdb2.Client
	if infOpts != nil {
		client = influxdb2.NewClientWithOptions(opt.Host, opt.Token, infOpts)
	} else {
		client = influxdb2.NewClient(opt.Host, opt.Token)
	}
	return &Client{
		Client: client,
	}
}
