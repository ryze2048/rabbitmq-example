package consumer

import (
	"context"
	"fmt"
	"github.com/ryze2048/rabbitmq-example/global"
	"github.com/ryze2048/rabbitmq-example/utils"
	"go.uber.org/zap"
	"time"
)

func Consumer(ctx context.Context) {
	global.ZAPLOG.Info(fmt.Sprintf("start consumer --> %v", time.Now().Format("2006-01-02 15:04:05")))

	var err error
	var rabbitmq *global.Rabbitmq
	if rabbitmq, err = utils.InitRabbitmq(global.ExchangeName, global.QueueName, global.RoutingKeyName, nil); err != nil {
		global.ZAPLOG.Error("init client queue err --> ", zap.Error(err))
		return
	}
	defer utils.CloseRabbitmqChannel(rabbitmq)

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
			if err = data.Acknowledger.Ack(data.DeliveryTag, false); err != nil {
				global.ZAPLOG.Error("ack err --> ", zap.Error(err))
			} else {
				global.ZAPLOG.Info(fmt.Sprintf("acknowledgement ack message --> %s", string(data.Body)))
			}
		}
	}
}
