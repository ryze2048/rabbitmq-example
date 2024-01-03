package utils

import (
	"github.com/ryze2048/rabbitmq-example/global"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

// InitRabbitmq 初始化信息
// exchangeName 交换机名称
// queueName 队列名
// routingKey 路由键
func InitRabbitmq(exchangeName, queueName, routingKey string, table amqp.Table) (*global.Rabbitmq, error) {
	var err error
	rabbitmq := InitRabbitmqMessage(exchangeName, queueName, routingKey, table)

	// 声明通道
	if rabbitmq.Channel, err = global.AMQP.Channel(); err != nil {
		global.ZAPLOG.Error("init amqp channel err --> ", zap.Error(err))
		return nil, err
	}

	//声明交换机
	if err = rabbitmq.Channel.ExchangeDeclare(exchangeName, amqp.ExchangeDirect, true, false, false, false, nil); err != nil {
		global.ZAPLOG.Error("init amqp exchange err --> ", zap.Error(err))
		return nil, err
	}

	// 声明队列信息
	if rabbitmq.Queue, err = rabbitmq.Channel.QueueDeclare(rabbitmq.QueueName, true, false, false, false, table); err != nil {
		global.ZAPLOG.Error("init amqp queue err --> ", zap.Error(err))
		return nil, err
	}

	// 绑定队列信息
	if err = rabbitmq.Channel.QueueBind(rabbitmq.Queue.Name, rabbitmq.RoutingKey, rabbitmq.ExchangeName, false, nil); err != nil {
		global.ZAPLOG.Error("init amqp queue bind err --> ", zap.Error(err))
		return nil, err
	}
	return rabbitmq, nil
}

// InitRabbitmqMessage InitMessage 信息绑定
func InitRabbitmqMessage(exchangeName, queueName, routingKey string, table amqp.Table) *global.Rabbitmq {
	return &global.Rabbitmq{
		ExchangeName: exchangeName,
		QueueName:    queueName,
		RoutingKey:   routingKey,
		Table:        table,
	}
}

// InitRabbitmqTable 绑定Table信息 - 死信队列
func InitRabbitmqTable(exchangeName, routingKey string) amqp.Table {
	return amqp.Table{
		"x-dead-letter-exchange":    exchangeName, // 指定死信交换机
		"x-dead-letter-routing-key": routingKey,   // 指定死信routing-key
	}
}

// CloseRabbitmqChannel 关闭通道
func CloseRabbitmqChannel(rabbitmq *global.Rabbitmq) {
	var err error
	if err = rabbitmq.Channel.Close(); err != nil {
		global.ZAPLOG.Error("close rabbitmq channel err --> ", zap.Error(err))
	}
}
