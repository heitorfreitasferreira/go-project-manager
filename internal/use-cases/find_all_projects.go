package usecases

import "github.com/heitorfreitasferreira/go-project-manager/internal/database"

type findAllProjects struct {
	projectRepository database.ProjectRepository
}

type FindAllProjectsOut []FindProjectByIdOut

type FindProjectsIn struct {
}

func (b *findAllProjects) Execute(in FindProjectsIn) (FindAllProjectsOut, error) {
	projetos, err := b.projectRepository.GetAllProject()
	if err != nil {
		return []FindProjectByIdOut{}, err
	}

	var projetosOut []FindProjectByIdOut
	for _, projeto := range projetos {
		projetosOut = append(projetosOut, FindProjectByIdOut{
			ID:          projeto.ID,
			Name:        projeto.Name.String,
			Description: projeto.Description.String,
			StartDate:   projeto.StartDate.Time,
			EndDate:     projeto.EndDate.Time,
		})
	}

	return projetosOut, nil
}
