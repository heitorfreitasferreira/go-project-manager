package usecases

import (
	"database/sql"
	"log"
	"time"

	"github.com/heitorfreitasferreira/go-project-manager/internal/database"
	"github.com/heitorfreitasferreira/go-project-manager/internal/models"
)

type createProject struct {
	projectRepository database.ProjectRepository
}

type CreateProjectIn struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Status      string    `json:"status"`
}

type CreateProjectOut struct {
	ID int
}

func (c *createProject) Execute(input CreateProjectIn) (CreateProjectOut, error) {
	projeto := &models.Project{
		Name:        sql.NullString{String: input.Name, Valid: true},
		Description: sql.NullString{String: input.Description, Valid: true},
		StartDate:   sql.NullTime{Time: input.StartDate, Valid: true},
		EndDate:     sql.NullTime{Time: input.EndDate, Valid: true},
		Status:      models.ToProjectStatus(input.Status),
	}

	err := c.projectRepository.CreateProject(projeto)
	if err != nil {
		log.Default().Printf("error creating project. Err: %v", err)
		return CreateProjectOut{}, err
	}

	return CreateProjectOut{ID: projeto.ID}, nil
}
