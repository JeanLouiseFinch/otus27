package main

import (
	"net"

	"github.com/JeanLouiseFinch/otus21/calendar"
	"github.com/JeanLouiseFinch/otus21/config"
	"github.com/JeanLouiseFinch/otus21/log"
	"github.com/JeanLouiseFinch/otus21/proto"

	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	cfg, err := config.GetConfig()
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
	reflection.Register(grpcServer)

	proto.RegisterCalendarServiceServer(grpcServer, calendar.NewServerCalendar(calendar.NewCalendar(l)))
	grpcServer.Serve(lis)
}
