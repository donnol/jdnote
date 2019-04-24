package message

import (
	"testing"

	utillog "github.com/donnol/jdnote/utils/log"
	nsq "github.com/nsqio/go-nsq"
)

func TestNsq(t *testing.T) {
	topic := "fruit"
	channel := "apple"
	config := nsq.NewConfig()

	// 生产消息
	nsqdAddr := "127.0.0.1:4150" // TCP: listening on [::]:4150
	producer, err := nsq.NewProducer(nsqdAddr, config)
	if err != nil {
		t.Fatal(err)
	}
	if err := producer.Publish(topic, []byte("Hello")); err != nil {
		t.Fatal(err)
	}

	// 消费消息
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		t.Fatal(err)
	}
	consumer.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		utillog.Debugf("%+v\n", m)
		return nil
	}))
	addr := "127.0.0.1:4161" // HTTP: listening on [::]:4161
	if err := consumer.ConnectToNSQLookupd(addr); err != nil {
		t.Fatal(err)
	}
}
