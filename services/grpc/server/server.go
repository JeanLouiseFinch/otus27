package main

import (
	"net"
	"os"
	"time"

	"github.com/JeanLouiseFinch/otus27/api/model"

	"github.com/JeanLouiseFinch/otus27/api/config"
	"github.com/JeanLouiseFinch/otus27/api/log"
	"github.com/JeanLouiseFinch/otus27/api/proto"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func main() {
	var (
		cfg *config.ConfigLog
		err error
	)
	if len(os.Args) > 1 {
		cfg, err = config.GetConfigLog(os.Args[1])
	} else {
		cfg, err = config.GetConfigLog("")
	}
	if err != nil {
		panic(err)
	}
	l, err := log.GetLogger(cfg.TypeLog)
	if err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Minute)
	l.Info("Running...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		l.Fatal("Failed to listen %v", zap.Error(err))
	}

	grpcServer := grpc.NewServer()

	cal, err := model.NewCalendar()
	if err != nil {
		l.Fatal("Not creating calendar", zap.Error(err))
	}
	proto.RegisterCalendarServiceServer(grpcServer, model.NewServerCalendar(cal, l))
	grpcServer.Serve(lis)
}
