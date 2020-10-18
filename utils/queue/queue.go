package queue

type Topic struct{}

type Func func()

type Param struct{}

type Option struct{}

type Trigger interface {
	Subscribe(topic Topic, f Func)
	Publish(topic Topic, param Param)
}

func NewTrigger(opt Option) Trigger {
	return &triggerImpl{}
}

// === 实现 ===

// TODO:
type triggerImpl struct {
}

// Publish 在本srv里通过topic告诉发布消息,说明自己发生了什么事,其它srv订阅这个topic,从而得到通知然后执行相应方法
// 如果你想别人关心你，你就发布消息告知别人
func (trigger *triggerImpl) Publish(topic Topic, param Param) {
	// 发到哪里，可以是queue

	// 失败了怎么办
}

// Subscribe 订阅其它srv的topic,在得到通知后执行自己的方法进行相应的处理
// 如果你关心别人，你就订阅Ta的消息
// 同一个Topic，可以绑定不同Func嘛？
func (trigger *triggerImpl) Subscribe(topic Topic, f Func) {

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
