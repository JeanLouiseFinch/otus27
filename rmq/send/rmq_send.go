package main

import (
	"fmt"
	"otus25/config"
	"otus25/internal/sql"
	"otus25/log"
	"time"

	"go.uber.org/zap"

	"github.com/streadway/amqp"
)

func main() {
	cfg, err := config.GetConfig("")
	if err != nil {
		panic(err)
	}
	l, err := log.GetLogger(cfg.TypeLog)
	if err != nil {
		panic(err)
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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
		"events", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		l.Fatal("Failed to declare a queue", zap.Error(err))
	}
	msg := ""
	for {
		events, err := sql.GetEventsByTime(5 * time.Minute)
		if err != nil {
			l.Error("Failed to get events", zap.Error(err))
			continue
		}
		for _, val := range events {
			msg = fmt.Sprintf("Event: %s (%s). Time: %v-%v", val.Title, val.Description, val.Start, val.End)
			err = ch.Publish(
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(msg),
				})
			if err != nil {
				l.Error("Failed to publish a message", zap.Error(err))
			} else {
				l.Info("Send", zap.String("event", val.Title))
			}
		}
		l.Info("Sleeping 1 minute")
		<-time.After(time.Minute * 1)
	}
}