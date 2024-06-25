package usecases

import (
	"database/sql"
	"raw-sqlite/internal/database"
	"raw-sqlite/internal/models"
	"time"
)

type createTask struct {
	taskRepository database.TaskRepository
}

type CreateTaskIn struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Owner       string             `json:"owner"`
	StartDate   *time.Time         `json:"start_date"`
	EndDate     *time.Time         `json:"end_date"`
	Status      *models.TaskStatus `json:"status"`
}

type CreateTaskOut struct{}

func (u createTask) Execute(in CreateTaskIn) (CreateTaskOut, error) {
	if in.StartDate == nil {
		in.StartDate = new(time.Time)
	}
	if in.Status == nil {
		*in.Status = models.NotStarted
	}

	err := u.taskRepository.CreateTask(&models.Task{
		Name:        sql.NullString{String: in.Name, Valid: true},
		Description: sql.NullString{String: in.Description, Valid: true},
		Owner:       sql.NullString{String: in.Owner, Valid: true},
		StartDate:   sql.NullTime{Time: *in.StartDate, Valid: in.StartDate != nil},
		EndDate:     sql.NullTime{Time: *in.EndDate, Valid: true},
		Status:      *in.Status,
	})
	if err != nil {
		return CreateTaskOut{}, err
	}
	return CreateTaskOut{}, nil
}
