package data

import (
	"time"

	"bitbucket.org/marcoboschetti/visionmonk/src/entities"
)

// InsertNewJob Adds a job in the repository
func InsertNewJob(job *entities.CalendarEvent) error {
	job.DateCreated = time.Now()
	job.DateCreated = time.Now()
	err := pgConnection.Insert(job)
	return err
}

// UpdateJob updates an existing job in the repository
func UpdateJob(job *entities.CalendarEvent) error {
	job.DateUpdated = time.Now()
	err := pgConnection.Update(job)
	return err
}

// UpdateJob updates an existing job in the repository
func GetCalendarEvent(sourceImgID string) (*entities.CalendarEvent, error) {

	job := &entities.CalendarEvent{}
	err := pgConnection.Model(job).WhereStruct(job).First()

	if err != nil {
		return nil, (err)
	}

	return job, err
}
