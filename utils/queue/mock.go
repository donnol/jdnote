package queue

import "context"

type TriggerMock struct {
	PublishFunc func(ctx context.Context, topic Topic, param Param)

	SubscribeFunc func(ctx context.Context, topic Topic, f Func)
}

var _ Trigger = &TriggerMock{}

func (mockRecv *TriggerMock) Publish(ctx context.Context, topic Topic, param Param) {
	mockRecv.PublishFunc(ctx, topic, param)
}

func (mockRecv *TriggerMock) Subscribe(ctx context.Context, topic Topic, f Func) {
	mockRecv.SubscribeFunc(ctx, topic, f)
}

type QueueMock struct {
	DequeueFunc func() (topic Topic, param Param)

	EnqueueFunc func(topic Topic, param Param)
}

var _ Queue = &QueueMock{}

func (mockRecv *QueueMock) Dequeue() (topic Topic, param Param) {
	return mockRecv.DequeueFunc()
}

func (mockRecv *QueueMock) Enqueue(topic Topic, param Param) {
	mockRecv.EnqueueFunc(topic, param)
}
