package usecases

import (
	"log"
	"raw-sqlite/internal/database"
	"raw-sqlite/internal/models"
	"time"
)

type createProject struct {
	ProjectRepository database.ProjetoRepository
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
	projeto := &models.Projeto{
		Nome:        input.Nome,
		Descricao:   input.Descricao,
		DataInicio:  input.DataInicio,
		DataTermino: input.DataTermino,
		Status:      models.ToProjectStatus(input.Status),
	}

	err := c.ProjectRepository.CreateProjeto(projeto)
	if err != nil {
		log.Default().Printf("error creating project. Err: %v", err)
		return CreateProjectOut{}, err
	}

	return CreateProjectOut{ID: projeto.ID}, nil
}
