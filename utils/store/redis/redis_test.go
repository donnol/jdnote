package redis

import (
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v4"
	"github.com/go-redis/redis/v7"
)

func TestRedis(t *testing.T) {
	gofakeit.Seed(time.Now().UnixNano())

	client := NewClient(&redis.Options{})

	testKey := "testKey"
	err := client.Set(testKey, gofakeit.Letter(), 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(testKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(testKey, val)

	testKey2 := "testKey2"
	val2, err := client.Get(testKey2).Result()
	if err == redis.Nil {
		fmt.Println(testKey2, " does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(testKey2, val2)
	}
}
