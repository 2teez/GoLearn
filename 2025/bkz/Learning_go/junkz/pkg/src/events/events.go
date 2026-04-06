package events

import (
	"errors"
	"fmt"
	"junkz/pkg/src/calendar"
	"unicode/utf8"
)

type Event struct {
	Title string
	*calendar.Date
}

func NewEvent(title string, date *calendar.Date) (*Event, error) {
	event := &Event{
		Title: title,
		Date:  date,
	}
	err := event.setTitle(title)
	if err != nil {
		return nil, err
	}
	return event, nil

}

func (e *Event) setTitle(title string) error {
	n := utf8.RuneCountInString(title)
	if n > 30 {
		return errors.New("Title can't be longer than 30 character letters.")
	}
	e.Title = title
	return nil
}

func (e *Event) String() string {
	return fmt.Sprintf("Event: %s on this date: %s", e.Title, e.Date)
}
