package models

import (
	"database/sql"
	"fmt"
)

type TaskStatus string

const (
	NotStarted TaskStatus = "NOT_STARTED"
	InProgress TaskStatus = "IN_PROGRESS"
	Completed  TaskStatus = "COMPLETED"
)

type Task struct {
	ID          int
	Name        sql.NullString
	Description sql.NullString
	Owner       sql.NullString
	StartDate   sql.NullTime
	EndDate     sql.NullTime
	Status      TaskStatus
	ProjectId   int
}

func (t *Task) String() string {
	return fmt.Sprintf("Task: ID: %d, Name: %s, Description: %s, Owner: %s, StartDate: %v, EndDate: %v, Status: %s, ProjectId: %d", t.ID, t.Name.String, t.Description.String, t.Owner.String, t.StartDate.Time, t.EndDate.Time, t.Status, t.ProjectId)
}
