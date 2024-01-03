package manual

// 手动确认
// 失败重试

type Manual struct {
}

const (
	ManualQueueName    string = `manual-queue`      // 队列名称
	ManualExchangeName string = `manual-exchange`   // 交换机名称
	ManualRoutingKey   string = `manual-routingKey` // 路由键
)

const (
	ManualDeadQueueName    string = `dead-manual-queue`      // 死信队列名称
	ManualDeadExchangeName string = `dead-manual-exchange`   // 死信交换机名称
	ManualDeadRoutingKey   string = `dead-manual-routingKey` // 死信路由键
)
