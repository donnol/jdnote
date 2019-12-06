package redis

import (
	"github.com/go-redis/redis/v7"
)

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
