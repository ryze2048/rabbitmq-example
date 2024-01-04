package delay

// 延迟队列
// 可以设置过期时间 防止设置延迟时间到了没有消费的情况
// 如果设置的延迟时间没有进行消费 会自动进入死信队列中 然后再进行消费

type Delay struct {
}

const (
	DelayQueueName    string = `delay-queue`      // 队列名称
	DelayExchangeName string = `delay-exchange`   // 交换机名称
	DelayRoutingKey   string = `delay-routingKey` // 路由键
)

const (
	DelayDeadQueueName    string = `dead-delay-queue`      // 队列名称
	DelayDeadExchangeName string = `dead-delay-exchange`   // 交换机名称
	DelayDeadRoutingKey   string = `dead-delay-routingKey` // 路由键
)
