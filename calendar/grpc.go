package calendar

import (
	"context"

	"go.uber.org/zap"

	"otus25/proto"

	"github.com/golang/protobuf/ptypes"
)

type ServerCalendar struct {
	calendar *Calendar
	logger   *zap.Logger
}

func NewServerCalendar(calendar *Calendar, logger *zap.Logger) *ServerCalendar {
	return &ServerCalendar{
		calendar: calendar,
		logger:   logger,
	}
}
func (c *ServerCalendar) NewEvent(context context.Context, in *proto.NewEventRequest) (*proto.NewEventResponse, error) {
	var err error
	event := &Event{}
	event.Description = in.GetEvent().GetDescription()
	event.Title = in.GetEvent().GetTitle()
	event.End, err = ptypes.Timestamp(in.GetEvent().GetEnd())
	if err != nil {
		c.logger.Error("Parse time start error", zap.Error(err))
		return nil, err
	}
	event.Start, err = ptypes.Timestamp(in.GetEvent().GetStart())
	if err != nil {
		c.logger.Error("Parse time start error", zap.Error(err))
		return nil, err
	}
	c.logger.Info("Inserting event ", zap.String("title", event.Title))
	uuid, err := c.calendar.NewEvent(event.Title, event.Description, event.Start, event.End)
	if err != nil {
		return nil, err
	}
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
	event.End, err = ptypes.Timestamp(in.GetEvent().GetEnd())
	if err != nil {
		c.logger.Error("Parse time start error", zap.Error(err))
		return nil, err
	}
	event.Start, err = ptypes.Timestamp(in.GetEvent().GetStart())
	if err != nil {
		c.logger.Error("Parse time start error", zap.Error(err))
		return nil, err
	}
	c.logger.Info("Modify event ", zap.String("title", event.Title))
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
	c.logger.Info("Remove event ", zap.String("id", id))
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
	c.logger.Info("Get event ", zap.String("id", id))
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
	resp.Event.End, err = ptypes.TimestampProto(event.End)
	if err != nil {
		c.logger.Error("Parse time end error", zap.Error(err))
		return nil, err
	}
	resp.Event.Start, err = ptypes.TimestampProto(event.Start)
	if err != nil {
		c.logger.Error("Parse time start error", zap.Error(err))
		return nil, err
	}
	return resp, err
}
