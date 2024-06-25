package models

import (
	"database/sql"
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
