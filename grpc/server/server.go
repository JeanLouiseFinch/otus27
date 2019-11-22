package main

import (
	"net"

	"github.com/JeanLouiseFinch/otus22/calendar"
	"github.com/JeanLouiseFinch/otus22/config"
	"github.com/JeanLouiseFinch/otus22/log"
	"github.com/JeanLouiseFinch/otus22/proto"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func main() {

	cfg, err := config.GetConfig("../../confita.yaml")
	if err != nil {
		panic(err)
	}
	l, err := log.GetLogger(cfg.TypeLog)
	if err != nil {
		panic(err)
	}
	l.Info("Running...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		l.Fatal("failed to listen %v", zap.Error(err))
	}

	grpcServer := grpc.NewServer()

	proto.RegisterCalendarServiceServer(grpcServer, calendar.NewServerCalendar(calendar.NewCalendar(), l))
	grpcServer.Serve(lis)
}
