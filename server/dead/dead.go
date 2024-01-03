package dead

import (
	"context"
	"fmt"
	"github.com/ryze2048/rabbitmq-example/global"
	"github.com/ryze2048/rabbitmq-example/utils"
	"go.uber.org/zap"
)

func Dead(ctx context.Context) {
	var err error
	var rabbitmq *global.Rabbitmq
	if rabbitmq, err = utils.InitRabbitmq(global.DeadExchangeName, global.DeadQueueName, global.DeadRoutingKeyName, nil); err != nil {
		global.ZAPLOG.Error("init  dead queue err --> ", zap.Error(err))
		return
	}
	defer utils.CloseRabbitmqChannel(rabbitmq)

	if err = rabbitmq.Channel.Qos(1, 0, false); err != nil {
		global.ZAPLOG.Error("channel qos err --> ", zap.Error(err))
		return
	}

	message, err := rabbitmq.Channel.Consume(rabbitmq.Queue.Name, "", false, false, false, false, nil)
	if err != nil {
		global.ZAPLOG.Error("init chapter consume err --> ", zap.Error(err))
		return
	}

	for {
		select {
		case info, ok := <-message:
			if !ok {
				global.ZAPLOG.Error("Failed to receive message from channel, exiting...")
				return
			}
			if err = info.Acknowledger.Ack(info.DeliveryTag, false); err != nil {
				global.ZAPLOG.Error("dead ack err --> ", zap.Error(err))
			} else {
				global.ZAPLOG.Info(fmt.Sprintf("dead acknowledgement ack message --> %s", string(info.Body)))
			}
		case <-ctx.Done():
			global.ZAPLOG.Info("收到停止通知")
			return
		}
	}

}
