package influx

import (
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type Client struct {
	influxdb2.Client
}

type Option struct {
	Host     string // 如：http://localhost:8086
	UserName string
	Password string
}

func makeToken(userName, password string) string {
	return fmt.Sprintf("%s:%s", userName, password)
}

func Open(opt Option, infOpts *influxdb2.Options) *Client {
	token := makeToken(opt.UserName, opt.Password)

	var client influxdb2.Client
	if infOpts != nil {
		client = influxdb2.NewClientWithOptions(opt.Host, token, infOpts)
	} else {
		client = influxdb2.NewClient(opt.Host, token)
	}
	return &Client{
		Client: client,
	}
}
