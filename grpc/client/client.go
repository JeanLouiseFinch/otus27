package main

import (
	"context"
	"fmt"

	"time"

	"otus22/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer cc.Close()

	c := proto.NewCalendarServiceClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()
	fmt.Printf("Insert:%s\n", "title 1")
	event1 := proto.Event{}
	event1.Title = "title 1"
	event1.Description = "description 1"
	event1.Start = ptypes.TimestampNow()
	event1.End = ptypes.TimestampNow()
	ev, err := c.NewEvent(ctx, &proto.NewEventRequest{Event: &event1})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Deadline exceeded!")
			} else {
				fmt.Printf("undexpected error %s\n", statusErr.Message())
			}
		} else {
			fmt.Printf("Error while calling RPC CheckHomework: %v", err)
		}
	} else {
		fmt.Println(ev.GetId())
	}
	fmt.Printf("Insert:%s\n", "title 2")
	event2 := proto.Event{}
	event2.Title = "title 2"
	event2.Description = "description 2"
	event2.Start = ptypes.TimestampNow()
	event2.End = ptypes.TimestampNow()
	ev, err = c.NewEvent(ctx, &proto.NewEventRequest{Event: &event2})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Deadline exceeded!")
			} else {
				fmt.Printf("undexpected error %s\n", statusErr.Message())
			}
		} else {
			fmt.Printf("Error while calling RPC CheckHomework: %v", err)
		}
	} else {
		fmt.Println(ev.GetId())
	}
	fmt.Printf("Insert:%s\n", "title 3")
	event3 := proto.Event{}
	event3.Title = "title 3"
	event3.Description = "description 3"
	event3.Start = ptypes.TimestampNow()
	event3.End = ptypes.TimestampNow()
	ev, err = c.NewEvent(ctx, &proto.NewEventRequest{Event: &event3})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Deadline exceeded!")
			} else {
				fmt.Printf("undexpected error %s\n", statusErr.Message())
			}
		} else {
			fmt.Printf("Error while calling RPC CheckHomework: %v", err)
		}
	} else {
		fmt.Println(ev.GetId())
		fmt.Println("end insert")
	}

	resp, err := c.GetEvent(ctx, &proto.GetEventRequest{Id: ev.GetId()})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Deadline exceeded!")
			} else {
				fmt.Printf("undexpected error %s\n", statusErr.Message())
			}
		} else {
			fmt.Printf("Error while calling RPC CheckHomework: %v", err)
		}
	} else {
		fmt.Printf("get %s by id %s\n", resp.GetEvent().GetTitle(), ev.GetId())
	}
	tt, _ := ptypes.TimestampProto(time.Now())
	_, err = c.ModifyEvent(ctx, &proto.ModifyEventRequest{Id: ev.GetId(), Event: &proto.Event{
		Title:       "title 3 modify",
		End:         tt,
		Start:       tt,
		Description: "descr 3 modify",
	}})

	resp, err = c.GetEvent(ctx, &proto.GetEventRequest{Id: ev.GetId()})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Deadline exceeded!")
			} else {
				fmt.Printf("undexpected error %s\n", statusErr.Message())
			}
		} else {
			fmt.Printf("Error while calling RPC CheckHomework: %v", err)
		}
	} else {
		fmt.Printf("get %s by id %s\n", resp.GetEvent().GetTitle(), ev.GetId())
	}
	_, err = c.RemoveEvent(ctx, &proto.RemoveEventRequest{Id: ev.GetId()})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Deadline exceeded!")
			} else {
				fmt.Printf("undexpected error %s\n", statusErr.Message())
			}
		} else {
			fmt.Printf("Error while calling RPC CheckHomework: %v", err)
		}
	} else {

	}
}
