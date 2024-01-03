package manual

import (
	"context"
	"fmt"
	"github.com/ryze2048/rabbitmq-example/global"
	"github.com/ryze2048/rabbitmq-example/process"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"time"
)

func (m *Manual) Dead(ctx context.Context) {
	global.ZAPLOG.Info(fmt.Sprintf("start dead  --> %v", time.Now().Format("2006-01-02 15:04:05")))

	var err error
	var rabbitmq *global.Rabbitmq
	if rabbitmq, err = process.InitRabbitmq(ManualDeadExchangeName, ManualDeadQueueName, ManualDeadRoutingKey, amqp.ExchangeDirect, nil); err != nil {
		global.ZAPLOG.Error("init client queue err --> ", zap.Error(err))
		return
	}
	defer process.CloseRabbitmqChannel(rabbitmq)

	if err = rabbitmq.Channel.Qos(1, 0, false); err != nil {
		global.ZAPLOG.Error("queue channel qos err --> ", zap.Error(err))
		return
	}

	message, err := rabbitmq.Channel.Consume(rabbitmq.Queue.Name, "", false, false, false, false, nil)
	if err != nil {
		global.ZAPLOG.Error("consume err --> ", zap.Error(err))
		return
	}

	for {
		select {
		case data, ok := <-message:
			if !ok {
				global.ZAPLOG.Error("Failed to receive message from channel, exiting...")
				return
			}
			time.Sleep(2 * time.Second)
			m.ack(data)
		case <-ctx.Done():
			global.ZAPLOG.Info("收到停止通知")
			return
		}
	}
}
