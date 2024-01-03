package delay

import (
	"context"
	"fmt"
	"github.com/ryze2048/rabbitmq-example/global"
	"github.com/ryze2048/rabbitmq-example/process"
	"go.uber.org/zap"
	"time"
)

func (d *Delay) Consumer(ctx context.Context) {
	global.ZAPLOG.Info(fmt.Sprintf("start automatic consumer --> %v", time.Now().Format("2006-01-02 15:04:05")))
	var err error
	var rabbitmq *global.Rabbitmq
	if rabbitmq, err = process.InitRabbitmq(DelayExchangeName, DelayQueueName, DelayRoutingKey, `x-delayed-message`, nil); err != nil {
		global.ZAPLOG.Error("init client queue err --> ", zap.Error(err))
		return
	}

	defer process.CloseRabbitmqChannel(rabbitmq)

	if err = rabbitmq.Channel.Qos(1, 0, false); err != nil {
		global.ZAPLOG.Error("queue channel qos err --> ", zap.Error(err))
		return
	}

	message, err := rabbitmq.Channel.Consume(rabbitmq.Queue.Name, "", true, false, false, false, nil)
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
			global.ZAPLOG.Info(fmt.Sprintf("consumer time --> %v, body info --> %s", time.Now().Format("2006-01-02 15:04:05"), string(data.Body)))
		case <-ctx.Done():
			global.ZAPLOG.Info("收到停止通知")
			return
		}
	}
}
