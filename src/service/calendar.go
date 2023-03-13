package service

import (
	"sync"

	"bitbucket.org/marcoboschetti/visionmonk/src/data"
	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
)

var calendarEventsUpdateMap = sync.Map{}

func GetCalendarEventsForToday() (*entities.CalendarEvent, error) {
	calendarEvent, err := data.GetCalendarEvent("asd")
	return calendarEvent, err
}

// PostNewJob insets a new calendarEvent in the DB and assigns a worker for it
func PostNewCalendarEvent(calendarEventID string) (*entities.CalendarEvent, error) {

	// Check cache
	calendarEvent, err := data.GetCalendarEvent(calendarEventID)
	return calendarEvent, err
}
