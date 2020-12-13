package redis

import (
	"github.com/go-redis/redis/v8"
)

type Options = redis.Options

type XAddArgs = redis.XAddArgs

type XReadArgs = redis.XReadArgs

type XReadGroupArgs = redis.XReadGroupArgs

// Client Client
type Client struct {
	*redis.Client
}

// NewClient NewClient
func NewClient(opt *redis.Options) *Client {
	c := &Client{}
	c.Client = redis.NewClient(opt)
	return c
}
