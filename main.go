package main

import (
	"context"
	"fmt"
	"github.com/ryze2048/rabbitmq-example/global"
	"github.com/ryze2048/rabbitmq-example/initialize"
	"github.com/ryze2048/rabbitmq-example/server/consumer"
	"github.com/ryze2048/rabbitmq-example/server/producer"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initialize.InitLog()
	initialize.InitAmqp()
	ctx, cancel := context.WithCancel(context.Background())
	go producer.Producer()
	go consumer.Consumer(ctx)
	// go dead.Dead(ctx)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case sig := <-sigChan:
			global.ZAPLOG.Info("To exit press CTRL+C")
			// 处理中断信号
			global.ZAPLOG.Info(fmt.Sprintf("Received signal: %s", sig))
			cancel()
			os.Exit(0)
		}
	}
}
