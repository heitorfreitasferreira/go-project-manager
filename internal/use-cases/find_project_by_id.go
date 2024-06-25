package usecases

import (
	"github.com/heitorfreitasferreira/go-project-manager/internal/database"
	"github.com/heitorfreitasferreira/go-project-manager/internal/models"
	"github.com/heitorfreitasferreira/go-project-manager/internal/types"
)

type findProjectByID struct {
	projectRepository database.ProjectRepository
}

type FindProjectByIdIn int

type FindProjectByIdOut types.Project

func (b *findProjectByID) Execute(in FindProjectByIdIn) (FindProjectByIdOut, error) {
	projeto, err := b.projectRepository.GetProjectoByID(int(in))
	if err != nil {
		return FindProjectByIdOut{}, err
	}

	var tasks []types.Task = make([]types.Task, len(projeto.Tasks))

	for i, task := range projeto.Tasks {
		tasks[i] = types.FromModelToTask(*task)
	}

	return FindProjectByIdOut{
		ID:          projeto.ID,
		Name:        projeto.Name.String,
		Description: projeto.Description.String,
		StartDate:   projeto.StartDate.Time,
		EndDate:     projeto.EndDate.Time,
		Status:      models.ProjectStatus(projeto.Status),
		Tasks:       tasks,
	}, nil
}
