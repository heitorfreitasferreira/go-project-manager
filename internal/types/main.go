package types

import (
	"time"

	"github.com/heitorfreitasferreira/go-project-manager/internal/models"
)

type Task struct {
	ID          int               `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Owner       string            `json:"owner"`
	StartDate   time.Time         `json:"start_date"`
	EndDate     time.Time         `json:"end_date"`
	Status      models.TaskStatus `json:"status"`
}

type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	StartDate time.Time            `json:"start_date"`
	EndDate   time.Time            `json:"end_date"`
	Status    models.ProjectStatus `json:"status"`
	Tasks     []Task               `json:"tasks"`
}

func FromModelToTask(in models.Task) Task {
	return Task{
		ID:          int(in.ID.Int64),
		Title:       in.Name.String,
		Description: in.Description.String,
		Owner:       in.Owner.String,
		StartDate:   in.StartDate.Time,
		EndDate:     in.EndDate.Time,
		Status:      in.Status,
	}
}

func FromModelToProject(in models.Project, tasks ...models.Task) Project {
	tasksDto := []Task{}
	for _, task := range tasks {
		tasksDto = append(tasksDto, FromModelToTask(task))
	}

	return Project{
		ID:          in.ID,
		Name:        in.Name.String,
		Description: in.Description.String,

		StartDate: in.StartDate.Time,
		EndDate:   in.EndDate.Time,
		Status:    in.Status,
		Tasks:     tasksDto,
	}
}
