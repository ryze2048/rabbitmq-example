package initialize

import (
	"github.com/ryze2048/rabbitmq-example/global"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func InitAmqp() {
	var err error
	if global.AMQP, err = amqp.Dial("amqp://root:123456@127.0.0.1:5672"); err != nil {
		global.ZAPLOG.Error("init amqp client err --> ", zap.Error(err))
		return
	}
	global.ZAPLOG.Info("init amqp client successful")
}
