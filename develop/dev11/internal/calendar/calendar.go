package calendar

import (
	"errors"
	"sync"
	"time"
)

type Calendar struct {
	events sync.Map
}

// Result is a structure for sending multiple events.
type Result struct {
	Result []Event `json:"result"`
}

// NewCalendar calendar constructor.
func NewCalendar() *Calendar {
	return &Calendar{}
}

// CreateEvent adds an event to the calendar.
func (c *Calendar) CreateEvent(event *Event) {
	c.events.Store(event.ID, *event)
}

// UpdateEvent updates an existing event.
func (c *Calendar) UpdateEvent(ID int, Title string, Date time.Time) error {
	value, ok := c.events.Load(ID)
	if !ok {
		return errors.New("no such event")
	}

	event := value.(Event) // Type casting.
	if !Date.IsZero() {
		event.Date = Date
	}
	if Title != "" {
		event.Title = Title
	}
	c.events.Store(ID, event)
	return nil
}

// DeleteEvent removes an event from the calendar.
func (c *Calendar) DeleteEvent(ID int) (*Event, error) {
	value, ok := c.events.LoadAndDelete(ID)
	if !ok {
		return nil, errors.New("no such event")
	}

	event := value.(Event)
	return &event, nil
}

// DailyEvents returns events for the current day.
func (c *Calendar) DailyEvents() []Event {
	var result []Event
	tYear, tMonth, tDay := time.Now().Date()

	c.events.Range(func(_, value interface{}) bool {
		event := value.(Event)
		year, month, day := event.Date.Date()
		if tYear == year && tMonth == month && tDay == day {
			result = append(result, event)
		}
		return true
	})

	return result
}

// WeeklyEvents returns events for the current week.
func (c *Calendar) WeeklyEvents() []Event {
	var result []Event
	tYear, tWeek := time.Now().ISOWeek()

	c.events.Range(func(_, value interface{}) bool {
		event := value.(Event)
		year, week := event.Date.ISOWeek()
		if tYear == year && tWeek == week {
			result = append(result, event)
		}
		return true
	})

	return result
}

// MonthlyEvents returns events for the current month.
func (c *Calendar) MonthlyEvents() []Event {
	var result []Event
	tYear, tMonth, _ := time.Now().Date()

	c.events.Range(func(_, value interface{}) bool {
		event := value.(Event)
		year, month, _ := event.Date.Date()
		if tYear == year && tMonth == month {
			result = append(result, event)
		}
		return true
	})

	return result
}
