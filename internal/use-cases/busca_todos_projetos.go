package usecases

import "raw-sqlite/internal/database"

type BuscaTodosProjetos struct {
	ProjectRepository database.ProjetoRepository
}

type BuscaTodosProjetosOut struct {
	Projetos []BuscaProjetoOut
}

type BuscaTodosProjetosIn struct {
}

func (b *BuscaTodosProjetos) Execute(in BuscaTodosProjetosIn) (BuscaTodosProjetosOut, error) {
	projetos, err := b.ProjectRepository.GetAllProjeto()
	if err != nil {
		return BuscaTodosProjetosOut{}, err
	}

	var projetosOut []BuscaProjetoOut
	for _, projeto := range projetos {
		projetosOut = append(projetosOut, BuscaProjetoOut{
			ID:          projeto.ID,
			Nome:        projeto.Nome,
			Descricao:   projeto.Descricao,
			DataInicio:  projeto.DataInicio,
			DataTermino: projeto.DataTermino,
		})
	}

	return BuscaTodosProjetosOut{Projetos: projetosOut}, nil
}
