package main

import (
	"context"
	"fmt"
	"time"

	"github.com/JeanLouiseFinch/otus27/api/proto"
	"github.com/cucumber/godog"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

const (
	queueName                 = "ToNotificationTest"
	notificationsExchangeName = "UserNotifications"
)

type newEventTest struct {
	clientConn *grpc.ClientConn
	calendar   proto.CalendarServiceClient

	returnError   error
	returnMessage *proto.NewEventResponse
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
func (net *newEventTest) start(interface{}) {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	panicOnErr(err)
	net.clientConn = cc
	net.calendar = proto.NewCalendarServiceClient(cc)
}

func (net *newEventTest) end(interface{}, error) {
	net.clientConn.Close()
}
func (net *newEventTest) iSendANewEventWith(arg1, arg2, arg3, arg4 string) error {
	event := proto.Event{}
	event.Title = arg1
	event.Description = arg2
	shortForm := "2006-Jan-02"
	start, _ := time.Parse(shortForm, arg3)
	start2, err := ptypes.TimestampProto(start)
	if err != nil {
		return err
	}
	event.Start = start2
	end, _ := time.Parse(shortForm, arg4)
	end2, err := ptypes.TimestampProto(end)
	if err != nil {
		return err
	}
	event.End = end2
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()
	res, err := net.calendar.NewEvent(ctx, &proto.NewEventRequest{Event: &event})
	net.returnError = err
	net.returnMessage = res
	return nil
}

func (net *newEventTest) theErrorShouldBeNil() error {
	if net.returnError != nil {
		return fmt.Errorf("unexpected error: %v != nil", net.returnError)
	}
	return nil
}
func (net *newEventTest) iSendISendANewHeader(arg1, arg2, arg3, arg4 string) error {
	event := proto.Event{}
	event.Title = arg1
	event.Description = arg2
	shortForm := "2006-Jan-02"
	start, _ := time.Parse(shortForm, arg3)
	start2, err := ptypes.TimestampProto(start)
	if err != nil {
		return err
	}
	event.Start = start2
	end, _ := time.Parse(shortForm, arg4)
	end2, err := ptypes.TimestampProto(end)
	if err != nil {
		return err
	}
	event.End = end2
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()
	_, err = net.calendar.ModifyEvent(ctx, &proto.ModifyEventRequest{
		Id:    net.returnMessage.GetId(),
		Event: &event,
	})
	net.returnError = err
	return nil
}

func (net *newEventTest) iSendTheEventId() error {
	time.Sleep(5 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()
	_, err := net.calendar.RemoveEvent(ctx, &proto.RemoveEventRequest{
		Id: net.returnMessage.GetId(),
	})
	net.returnError = err
	return nil
}

func FeatureContext(s *godog.Suite) {
	net := new(newEventTest)
	s.BeforeScenario(net.start)
	s.Step(`^I send a new event with "([^"]*)", "([^"]*)", "([^"]*)", "([^"]*)"$`, net.iSendANewEventWith)
	s.Step(`^the error should be nil$`, net.theErrorShouldBeNil)
	s.Step(`^I send I send a new header "([^"]*)", "([^"]*)", "([^"]*)", "([^"]*)"$`, net.iSendISendANewHeader)
	s.Step(`^the error should be nil$`, net.theErrorShouldBeNil)
	s.Step(`^I send the event id$`, net.iSendTheEventId)
	s.Step(`^the error should be nil$`, net.theErrorShouldBeNil)
	s.AfterScenario(net.end)
}
