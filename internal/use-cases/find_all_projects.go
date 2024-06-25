package usecases

import (
	"log"

	"github.com/heitorfreitasferreira/go-project-manager/internal/database"
	"github.com/heitorfreitasferreira/go-project-manager/internal/types"
)

type findAllProjects struct {
	projectRepository database.ProjectRepository
}

type FindAllProjectsOut []types.Project

type FindProjectsIn struct {
}

func (b *findAllProjects) Execute(in FindProjectsIn) (FindAllProjectsOut, error) {
	projetos, err := b.projectRepository.GetAllProject()
	if err != nil {
		return []types.Project{}, err
	}

	var projetosOut []types.Project
	for _, projeto := range projetos {
		var tasksOut []types.Task
		for _, task := range projeto.Tasks {
			tasksOut = append(tasksOut, types.Task{
				ID:          task.ID,
				Title:       task.Name.String,
				Description: task.Description.String,
				StartDate:   task.StartDate.Time,
				EndDate:     task.EndDate.Time,
			})
		}

		log.Default().Printf("Number of tasks: %d", len(tasksOut))
		projetosOut = append(projetosOut, types.Project{
			ID:          projeto.ID,
			Name:        projeto.Name.String,
			Description: projeto.Description.String,
			StartDate:   projeto.StartDate.Time,
			EndDate:     projeto.EndDate.Time,
			Tasks:       tasksOut,
		})
	}

	return projetosOut, nil
}
