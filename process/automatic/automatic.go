package automatic

type Automatic struct{}

const (
	AutomaticQueueName    string = `automatic-queue`      // 队列名称
	AutomaticExchangeName string = `automatic-exchange`   // 交换机名称
	AutomaticRoutingKey   string = `automatic-routingKey` // 路由键
)
