package main

import (
	"github.com/ryze2048/rabbitmq-example/initialize"
)

func main() {
	initialize.InitLog()
	initialize.InitAmqp()
}
