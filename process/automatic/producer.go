package automatic

import (
	"fmt"
	"github.com/ryze2048/rabbitmq-example/global"
	"github.com/ryze2048/rabbitmq-example/process"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"time"
)

// 自动确认机制
// 生产者

func (a *Automatic) Producer() {
	global.ZAPLOG.Info(fmt.Sprintf("start automatic producer --> %v", time.Now().Format("2006-01-02 15:04:05")))
	var err error
	var rabbitmq *global.Rabbitmq
	if rabbitmq, err = process.InitRabbitmq(AutomaticExchangeName, AutomaticQueueName, AutomaticRoutingKey, amqp.ExchangeDirect, nil); err != nil {
		global.ZAPLOG.Error("init client queue err --> ", zap.Error(err))
		return
	}

	defer process.CloseRabbitmqChannel(rabbitmq)

	for i := 0; i <= 10; i++ {
		err = rabbitmq.Channel.Publish(rabbitmq.ExchangeName, rabbitmq.RoutingKey, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05"))),
		})
		if err != nil {
			global.ZAPLOG.Error("push err --> ", zap.Error(err))
		}
		time.Sleep(1 * time.Second)
	}

}
