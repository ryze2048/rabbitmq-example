package initialize

import (
	"github.com/ryze2048/rabbitmq-example/global"
	"go.uber.org/zap"
)

func InitLog() {
	// logger := zap.NewExample()
	global.ZAPLOG = zap.NewExample()
}
