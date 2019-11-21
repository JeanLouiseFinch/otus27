package calendar

import (
	"context"

	"github.com/JeanLouiseFinch/otus21/proto"

	"github.com/golang/protobuf/ptypes"
	"go.uber.org/zap"
)

type ServerCalendar struct {
	calendar *Calendar
}

func NewServerCalendar(calendar *Calendar) *ServerCalendar {
	return &ServerCalendar{
		calendar: calendar,
	}
}
func (c *ServerCalendar) NewEvent(context context.Context, in *proto.NewEventRequest) (*proto.NewEventResponse, error) {
	var err error
	event := &Event{}
	event.Description = in.GetEvent().GetDescription()
	event.Title = in.GetEvent().GetTitle()
	event.Duration, err = ptypes.Duration(in.GetEvent().GetDuration())
	if err != nil {
		c.calendar.logger.Error("Parse duration error", zap.Error(err))
		return nil, err
	}
	event.Start, err = ptypes.Timestamp(in.GetEvent().GetStart())
	if err != nil {
		c.calendar.logger.Error("Parse time start error", zap.Error(err))
		return nil, err
	}
	uuid := c.calendar.NewEvent(event.Title, event.Description, event.Start, event.Duration)
	return &proto.NewEventResponse{Id: uuid}, nil
}
func (c *ServerCalendar) ModifyEvent(ctx context.Context, in *proto.ModifyEventRequest) (*proto.ModifyEventResponse, error) {
	var (
		event Event
		err   error
		id    string
	)
	event = Event{}
	id = in.GetId()
	event.Description = in.GetEvent().GetDescription()
	event.Title = in.GetEvent().GetTitle()
	event.Duration, err = ptypes.Duration(in.GetEvent().GetDuration())
	if err != nil {
		c.calendar.logger.Error("Parse duration error", zap.Error(err))
		return nil, err
	}
	event.Start, err = ptypes.Timestamp(in.GetEvent().GetStart())
	if err != nil {
		c.calendar.logger.Error("Parse time start error", zap.Error(err))
		return nil, err
	}
	err = c.calendar.ModifyEvent(id, event)
	return &proto.ModifyEventResponse{}, err
}
func (c *ServerCalendar) RemoveEvent(ctx context.Context, in *proto.RemoveEventRequest) (*proto.RemoveEventResponse, error) {
	var (
		err error
		id  string
	)
	id = in.GetId()
	err = c.calendar.RemoveEvent(id)
	if err != nil {
		return &proto.RemoveEventResponse{Ok: false}, err
	}
	return &proto.RemoveEventResponse{Ok: true}, err
}
func (c *ServerCalendar) GetEvent(ctx context.Context, in *proto.GetEventRequest) (*proto.GetEventResponse, error) {
	var (
		err   error
		id    string
		event Event
		resp  *proto.GetEventResponse
	)
	id = in.GetId()

	event, err = c.calendar.GetEvent(id)
	if err != nil {
		return &proto.GetEventResponse{Ok: false}, err
	}
	resp = &proto.GetEventResponse{
		Event: &proto.Event{
			Description: event.Description,
			Title:       event.Title,
		},
		Ok: true,
	}
	resp.Event.Start, err = ptypes.TimestampProto(event.Start)
	if err != nil {
		c.calendar.logger.Error("Parse time start error", zap.Error(err))
		return nil, err
	}
	resp.Event.Duration = ptypes.DurationProto(event.Duration)
	return resp, err
}
