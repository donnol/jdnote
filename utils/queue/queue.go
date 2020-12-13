package queue

import (
	"context"
	"fmt"

	"github.com/donnol/jdnote/utils/store/redis"
)

type Topic struct{}

type Func func(Param)

type Param struct {
	Values map[string]interface{}
}

type Option struct {
	RedisClient *redis.Client
}

type Trigger interface {
	Subscribe(ctx context.Context, topic Topic, f Func)
	Publish(ctx context.Context, topic Topic, param Param)
}

func NewTrigger(opt Option) Trigger {
	return &triggerImpl{
		redisClient: opt.RedisClient,
	}
}

// === 实现 ===

// TODO: 基于redis stream实现
type triggerImpl struct {
	redisClient *redis.Client
}

// Publish 在本srv里通过topic告诉发布消息,说明自己发生了什么事,其它srv订阅这个topic,从而得到通知然后执行相应方法
// 如果你想别人关心你，你就发布消息告知别人
func (trigger *triggerImpl) Publish(ctx context.Context, topic Topic, param Param) {
	// 发到哪里，可以是queue
	r, err := trigger.redisClient.XAdd(ctx, &redis.XAddArgs{
		Stream:       "mystream",
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values:       param.Values,
	}).Result()
	if err != nil {
		fmt.Printf("Publish failed, err: %+v\n", err)
	} else {
		fmt.Printf("Publish success, r: %+v\n", r)
	}

	// 失败了怎么办
}

// Subscribe 订阅其它srv的topic,在得到通知后执行自己的方法进行相应的处理
// 如果你关心别人，你就订阅Ta的消息
// 同一个Topic，可以绑定不同Func嘛？
func (trigger *triggerImpl) Subscribe(ctx context.Context, topic Topic, f Func) {
	r, err := trigger.redisClient.XRead(ctx, &redis.XReadArgs{
		Streams: []string{"mystream", "$"}, // list of streams and ids, e.g. stream1 stream2 id1 id2
		Count:   1,
		Block:   0,
	}).Result()
	if err != nil {
		fmt.Printf("Subscribe failed, err: %+v\n", err)
	} else {
		fmt.Printf("Subscribe success, r: %+v\n", r)
	}

	if len(r) != 0 {
		f(Param{
			Values: r[0].Messages[0].Values,
		})
	}
}

var (
	_ Trigger = &triggerImpl{}
)

type Queue interface {
	// 入队操作
	Enqueue(topic Topic, param Param)
	// 出队操作
	Dequeue() (topic Topic, param Param)
}
