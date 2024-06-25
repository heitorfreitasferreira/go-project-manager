package usecases

import "raw-sqlite/internal/database"

type buscaTodosProjetos struct {
	projectRepository database.ProjectRepository
}

type BuscaTodosProjetosOut []BuscaProjetoOut

type BuscaTodosProjetosIn struct {
}

func (b *buscaTodosProjetos) Execute(in BuscaTodosProjetosIn) (BuscaTodosProjetosOut, error) {
	projetos, err := b.projectRepository.GetAllProject()
	if err != nil {
		return []BuscaProjetoOut{}, err
	}

	var projetosOut []BuscaProjetoOut
	for _, projeto := range projetos {
		projetosOut = append(projetosOut, BuscaProjetoOut{
			ID:          projeto.ID,
			Nome:        projeto.Name.String,
			Descricao:   projeto.Description.String,
			DataInicio:  projeto.StartDate.Time,
			DataTermino: projeto.EndDate.Time,
		})
	}

	return projetosOut, nil
}
