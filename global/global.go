package global

import (
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var ZAPLOG *zap.Logger

var AMQP *amqp.Connection
