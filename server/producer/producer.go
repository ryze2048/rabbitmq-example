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
	if rabbitmq, err = utils.InitRabbitmq(global.ExchangeName, global.QueueName, global.RoutingKeyName, nil); err != nil {
		// if rabbitmq, err = utils.InitRabbitmq(global.ExchangeName, global.QueueName, global.RoutingKeyName, utils.InitRabbitmqTable(global.DeadExchangeName, global.DeadRoutingKeyName)); err != nil {
		global.ZAPLOG.Error("init client queue err --> ", zap.Error(err))
		return
	}
	defer utils.CloseRabbitmqChannel(rabbitmq)
	err = rabbitmq.Channel.Publish(rabbitmq.ExchangeName, rabbitmq.RoutingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Headers:     amqp.Table{"x-delay": 5000}, // 设置延迟时间，单位为毫秒
		Body:        []byte(fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05"))),
	})
	if err != nil {
		global.ZAPLOG.Error("push err --> ", zap.Error(err))
	}
}
