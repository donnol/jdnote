package queue

import (
	"context"
	"fmt"
	"testing"

	"github.com/donnol/jdnote/utils/store/redis"
)

func TestQueue(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "jdis1gHR",
	})
	trigger := NewTrigger(Option{
		RedisClient: client,
	})

	ctx := context.Background()

	ch1 := make(chan struct{})
	go func() {
		trigger.Subscribe(ctx, Topic{}, Func(func(p Param) {
			fmt.Printf("Subscribe: %+v\n", p)
		}))

		ch1 <- struct{}{}
	}()

	ch2 := make(chan struct{})
	go func() {
		trigger.Subscribe(ctx, Topic{}, Func(func(p Param) {
			fmt.Printf("Subscribe: %+v\n", p)
		}))

		ch2 <- struct{}{}
	}()

	trigger.Publish(ctx, Topic{}, Param{
		Values: map[string]interface{}{
			"name": "jd",
		},
	})

	<-ch1
	<-ch2

	// 怎么把在发布之后把之前的消息追回来呢？
	// trigger.Subscribe(ctx, Topic{}, Func(func() {
	// 	fmt.Printf("Subscribe\n")
	// }))

}
