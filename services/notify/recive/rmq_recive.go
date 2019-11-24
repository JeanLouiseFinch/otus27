package main

import (
	"fmt"
	"os"

	"github.com/JeanLouiseFinch/otus27/api/config"
	"github.com/JeanLouiseFinch/otus27/api/log"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func main() {
	var (
		cfg *config.Config
		err error
	)
	if len(os.Args) > 1 {
		cfg, err = config.GetConfig(os.Args[1])
	} else {
		cfg, err = config.GetConfig("")
	}

	if err != nil {
		panic(err)
	}
	l, err := log.GetLogger(cfg.Log.TypeLog)
	if err != nil {
		panic(err)
	}

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.RMQ.UserRMQ, cfg.RMQ.PasswordRMQ, cfg.RMQ.HostRMQ, cfg.RMQ.PortRMQ))
	if err != nil {
		l.Fatal("Failed to connect to RabbitMQ", zap.Error(err))
	}
	defer conn.Close()
	l.Info("Connected to RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		l.Fatal("Failed to open a channel", zap.Error(err))
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		cfg.RMQ.QueueRMQ, // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		l.Fatal("Failed to declare a queue", zap.Error(err))
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		l.Fatal("Failed to register a consumer", zap.Error(err))
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			l.Info("Recive", zap.String("message", fmt.Sprintf("%s", d.Body)))
		}
	}()

	l.Info("[*] Waiting for messages. To exit press CTRL+C")

	<-forever
}
