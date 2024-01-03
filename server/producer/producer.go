package producer

import (
	"fmt"
	"github.com/ryze2048/rabbitmq-example/global"
	"github.com/ryze2048/rabbitmq-example/utils"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"time"
)

func Producer() {
	global.ZAPLOG.Info(fmt.Sprintf("start producer --> %v", time.Now().Format("2006-01-02 15:04:05")))

	var err error
	var rabbitmq *global.Rabbitmq
	if rabbitmq, err = utils.InitRabbitmq("info-exchange", "info-queue", "info-routingKey", nil); err != nil {
		global.ZAPLOG.Error("init client queue err --> ", zap.Error(err))
		return
	}
	defer func() {
		if err = rabbitmq.Channel.Close(); err != nil {
			global.ZAPLOG.Error("close client channel err --> ", zap.Error(err))
		}
	}()
	err = rabbitmq.Channel.Publish(rabbitmq.ExchangeName, rabbitmq.RoutingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05"))),
	})
	if err != nil {
		global.ZAPLOG.Error("push err --> ", zap.Error(err))
	}
}
