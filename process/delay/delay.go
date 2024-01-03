package delay

type Delay struct {
}

const (
	DelayQueueName    string = `delay-queue`      // 队列名称
	DelayExchangeName string = `delay-exchange`   // 交换机名称
	DelayRoutingKey   string = `delay-routingKey` // 路由键
)
