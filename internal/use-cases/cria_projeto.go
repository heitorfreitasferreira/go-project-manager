package usecases

import (
	"database/sql"
	"log"
	"raw-sqlite/internal/database"
	"raw-sqlite/internal/models"
	"time"
)

type createProject struct {
	projectRepository database.ProjectRepository
}

type CreateProjectIn struct {
	Nome        string
	Descricao   string
	DataInicio  time.Time
	DataTermino time.Time
	Status      string
}

type CreateProjectOut struct {
	ID int
}

func (c *createProject) Execute(input CreateProjectIn) (CreateProjectOut, error) {
	projeto := &models.Project{
		Nome:        sql.NullString{String: input.Nome, Valid: true},
		Descricao:   sql.NullString{String: input.Descricao, Valid: true},
		DataInicio:  sql.NullTime{Time: input.DataInicio, Valid: true},
		DataTermino: sql.NullTime{Time: input.DataTermino, Valid: true},
		Status:      models.ToProjectStatus(input.Status),
	}

	err := c.projectRepository.CreateProject(projeto)
	if err != nil {
		log.Default().Printf("error creating project. Err: %v", err)
		return CreateProjectOut{}, err
	}

	return CreateProjectOut{ID: projeto.ID}, nil
}
