package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/JeanLouiseFinch/otus27/api/config"
	"github.com/JeanLouiseFinch/otus27/api/log"
	"github.com/JeanLouiseFinch/otus27/api/model"

	"go.uber.org/zap"

	"github.com/streadway/amqp"
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
	msg := ""
	calendar, err := model.NewCalendar()
	if err != nil {
		l.Fatal("Failed to create a calendar", zap.Error(err))
	}
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		events, err := calendar.GetEventsByTime(ctx, time.Duration(cfg.RMQ.DurationRMQ)*time.Minute)
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
		l.Info("Sleeping", zap.Int("minute", cfg.RMQ.TimeoutRMQ))
		<-time.After(time.Minute * time.Duration(cfg.RMQ.TimeoutRMQ))
	}
}
