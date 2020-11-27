package queue

type TriggerMock struct {
	PublishFunc func(topic Topic, param Param)

	SubscribeFunc func(topic Topic, f Func)
}

var _ Trigger = &TriggerMock{}

func (mockRecv *TriggerMock) Publish(topic Topic, param Param) {
	mockRecv.PublishFunc(topic, param)
}

func (mockRecv *TriggerMock) Subscribe(topic Topic, f Func) {
	mockRecv.SubscribeFunc(topic, f)
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
