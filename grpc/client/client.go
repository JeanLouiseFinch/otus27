package main

import (
	"context"
	"fmt"

	"time"

	"github.com/JeanLouiseFinch/otus25/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

func newEvent(c proto.CalendarServiceClient, title, descr string, start, end time.Time) error {
	event := proto.Event{}
	event.Title = title
	event.Description = descr
	start2, err := ptypes.TimestampProto(start)
	if err != nil {
		return err
	}
	event.Start = start2
	end2, err := ptypes.TimestampProto(end)
	if err != nil {
		return err
	}
	event.End = end2
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()
	_, err = c.NewEvent(ctx, &proto.NewEventRequest{Event: &event})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				return fmt.Errorf("Deadline exceeded! %v", err)
			} else {
				return fmt.Errorf("undexpected error %s\n", statusErr.Message())
			}
		} else {
			return fmt.Errorf("Error while calling RPC NewEvent: %v", err)
		}
	}
	return nil
}
func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer cc.Close()

	c := proto.NewCalendarServiceClient(cc)
	err = newEvent(c, "rmq1", "rmq1 descr", time.Now().Add(1*time.Minute), time.Now().Add(3*time.Minute))
	if err != nil {
		fmt.Println(err)
	}
	err = newEvent(c, "rmq2", "rmq2 descr", time.Now().Add(2*time.Minute), time.Now().Add(3*time.Minute))
	if err != nil {
		fmt.Println(err)
	}
	err = newEvent(c, "rmq3", "rmq3 descr", time.Now().Add(2*time.Minute), time.Now().Add(3*time.Minute))
	if err != nil {
		fmt.Println(err)
	}
	err = newEvent(c, "rmq4", "rmq4 descr", time.Now().Add(5*time.Minute), time.Now().Add(14*time.Minute))
	if err != nil {
		fmt.Println(err)
	}
	err = newEvent(c, "rmq5", "rmq5 descr", time.Now().Add(6*time.Minute), time.Now().Add(7*time.Minute))
	if err != nil {
		fmt.Println(err)
	}
	err = newEvent(c, "rmq6", "rmq6 descr", time.Now().Add(8*time.Minute), time.Now().Add(10*time.Minute))
	if err != nil {
		fmt.Println(err)
	}
	err = newEvent(c, "rmq7", "rmq7 descr", time.Now().Add(10*time.Minute), time.Now().Add(12*time.Minute))
	if err != nil {
		fmt.Println(err)
	}
}
