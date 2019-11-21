package calendar

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/google/uuid"
)

type Calendar struct {
	events map[string]*Event
	mutex  *sync.Mutex
	logger *zap.Logger
}

type Event struct {
	Title       string
	Description string
	Start       time.Time
	Duration    time.Duration
}

func NewCalendar(logger *zap.Logger) *Calendar {
	c := &Calendar{
		mutex:  &sync.Mutex{},
		events: make(map[string]*Event),
		logger: logger,
	}
	c.logger.Info("Calendar is creating...")
	return c
}

func (c *Calendar) NewEvent(title, description string, start time.Time, duration time.Duration) string {
	id := uuid.New().String()
	c.mutex.Lock()
	c.events[id] = &Event{
		Title:       title,
		Description: description,
		Start:       start,
		Duration:    duration,
	}
	c.logger.Info("New event", zap.String("title", title))
	c.mutex.Unlock()
	return id
}
func (c *Calendar) String() string {
	var result string
	c.mutex.Lock()
	result = fmt.Sprintf("Calendar:\n---\n")
	for key, val := range c.events {
		result += fmt.Sprintf("\tEvent->\n\t\tID: %v\n\t\tTitle: %s\n\t\tDescription: %s\n\t\tStart: %v\n\t\tDuration: %v\n", key, val.Title, val.Description, val.Start, val.Duration)
	}
	result += fmt.Sprintf("---\n")
	c.mutex.Unlock()
	return result
}

func (c *Calendar) GetEvent(id string) (Event, error) {
	if val, ok := c.events[id]; !ok {
		return Event{}, errors.New("Not event in database")
	} else {
		return *val, nil
	}
}

func (c *Calendar) ModifyEvent(id string, e Event) error {
	if _, ok := c.events[id]; !ok {
		return errors.New("Not event in database")
	} else {
		c.mutex.Lock()
		c.events[id].Description = e.Description
		c.events[id].Title = e.Title
		c.events[id].Start = e.Start
		c.events[id].Duration = e.Duration
		c.logger.Info("Modify event", zap.String("title", c.events[id].Title))
		c.mutex.Unlock()
		return nil
	}
}

func (c *Calendar) RemoveEvent(id string) error {
	if _, ok := c.events[id]; !ok {
		return errors.New("Not event in database")
	} else {
		c.mutex.Lock()
		c.logger.Info("Remove event", zap.String("title", c.events[id].Title))
		delete(c.events, id)
		c.mutex.Unlock()
		return nil
	}
}
