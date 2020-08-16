package queue

// 主题定义
const (
	TopicOne = "One"
)

// Base 基底
type Base struct {
}

// Publish 发布
func (b *Base) Publish(topic string, body []byte) error {
	return nil
}

// Subscribe 订阅
func (b *Base) Subscribe(topic, channel string) {

}

// Produce 生产
func (b *Base) Produce(topic string, body []byte) error {
	return nil
}

// Consume 消费
func (b *Base) Consume(topic, channel string) {

}
