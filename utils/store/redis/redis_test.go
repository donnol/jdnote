package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v4"
	"github.com/go-redis/redis/v8"
)

func TestRedis(t *testing.T) {
	ctx := context.Background()
	gofakeit.Seed(time.Now().UnixNano())

	client := NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	testKey := "testKey"
	err := client.Set(ctx, testKey, gofakeit.Letter(), 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, testKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(testKey, val)

	testKey2 := "testKey2"
	val2, err := client.Get(ctx, testKey2).Result()
	if err == redis.Nil {
		fmt.Println(testKey2, " does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(testKey2, val2)
	}
}
