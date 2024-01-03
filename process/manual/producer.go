package manual

import (
	"fmt"
	"github.com/ryze2048/rabbitmq-example/global"
	"github.com/ryze2048/rabbitmq-example/process"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func (m *Manual) Producer() {
	global.ZAPLOG.Info(fmt.Sprintf("start manual producer --> %v", time.Now().Format("2006-01-02 15:04:05")))

	var err error
	var rabbitmq *global.Rabbitmq
	if rabbitmq, err = process.InitRabbitmq(ManualExchangeName, ManualQueueName, ManualRoutingKey, amqp.ExchangeDirect, process.InitRabbitmqTable(ManualDeadExchangeName, ManualDeadRoutingKey)); err != nil {
		global.ZAPLOG.Error("init client queue err --> ", zap.Error(err))
		return
	}
	defer process.CloseRabbitmqChannel(rabbitmq)

	for i := 0; i <= 10; i++ {
		err = rabbitmq.Channel.Publish(rabbitmq.ExchangeName, rabbitmq.RoutingKey, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(strconv.Itoa(i)),
		})
		if err != nil {
			global.ZAPLOG.Error("push err --> ", zap.Error(err))
		}
	}
}
