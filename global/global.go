package global

import (
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var ZAPLOG *zap.Logger

var AMQP *amqp.Connection

type Rabbitmq struct {
	Channel      *amqp.Channel `json:"channel"`       // 通道
	Queue        amqp.Queue    `json:"queue"`         // 队列
	Table        amqp.Table    `json:"table"`         // 绑定死信交换机
	ExchangeName string        `json:"exchange_name"` // 交换机名称
	QueueName    string        `json:"queue_name"`    // 队列名称
	RoutingKey   string        `json:"routing_key"`   // 路由键
}
