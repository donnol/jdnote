package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/donnol/jdnote/utils/store/redis"
)

type Cache interface {
	Get(ctx context.Context, key string) (value interface{}, err error)
	Set(ctx context.Context, key string, value interface{}, exp time.Duration) (err error)
}

type Option struct {
	RedisClient *redis.Client
}

func New(opt Option) Cache {
	return &redisCacheImpl{
		client: opt.RedisClient,
	}
}

// redisCacheImpl 基于redis实现
type redisCacheImpl struct {
	client *redis.Client
}

func (impl *redisCacheImpl) logPrefix() string {
	return "| redisCacheImpl |"
}

func (impl *redisCacheImpl) Get(ctx context.Context, key string) (value interface{}, err error) {
	if value, err = impl.client.Get(ctx, key).Result(); err != nil {
		return nil, fmt.Errorf(impl.logPrefix()+"get failed, err: %w", err)
	}
	return
}

func (impl *redisCacheImpl) Set(ctx context.Context, key string, value interface{}, exp time.Duration) (err error) {
	if err = impl.client.Set(ctx, key, value, exp).Err(); err != nil {
		return fmt.Errorf(impl.logPrefix()+"set failed, err: %w", err)
	}
	return
}
