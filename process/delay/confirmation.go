package delay

import (
	"fmt"
	"github.com/ryze2048/rabbitmq-example/global"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func (m *Delay) nack(data amqp.Delivery) {
	var err error
	if err = data.Acknowledger.Nack(data.DeliveryTag, false, false); err != nil {
		global.ZAPLOG.Error("acknowledgement nack err --> ", zap.Error(err))
	}
	global.ZAPLOG.Info(fmt.Sprintf("acknowledgement nack message --> %s", string(data.Body)))
}

func (c *Delay) ack(data amqp.Delivery) {
	var err error
	if err = data.Acknowledger.Ack(data.DeliveryTag, false); err != nil {
		global.ZAPLOG.Error("acknowledgement ack err --> ", zap.Error(err))
	} else {
		global.ZAPLOG.Info(fmt.Sprintf("acknowledgement ack message --> %s", string(data.Body)))
	}
}
