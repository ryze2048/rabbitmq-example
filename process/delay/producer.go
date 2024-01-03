package delay

import (
	"fmt"
	"github.com/ryze2048/rabbitmq-example/global"
	"github.com/ryze2048/rabbitmq-example/process"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func (d *Delay) Producer() {

	global.ZAPLOG.Info(fmt.Sprintf("start automatic producer --> %v", time.Now().Format("2006-01-02 15:04:05")))
	var err error
	var rabbitmq *global.Rabbitmq
	if rabbitmq, err = process.InitRabbitmq(DelayExchangeName, DelayQueueName, DelayRoutingKey, `x-delayed-message`, nil); err != nil {
		global.ZAPLOG.Error("init client queue err --> ", zap.Error(err))
		return
	}

	defer process.CloseRabbitmqChannel(rabbitmq)

	for i := 0; i <= 10; i++ {
		err = rabbitmq.Channel.Publish(rabbitmq.ExchangeName, rabbitmq.RoutingKey, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Headers:     amqp.Table{"x-delay": 5000}, // 设置延迟时间，单位为毫秒
			Body:        []byte(strconv.Itoa(i)),
		})

		global.ZAPLOG.Info(fmt.Sprintf("producer time --> %v, body info --> %s", time.Now().Format("2006-01-02 15:04:05"), strconv.Itoa(i)))

		if err != nil {
			global.ZAPLOG.Error("producer err --> ", zap.Error(err))
		}
		time.Sleep(1 * time.Second)
	}

}
