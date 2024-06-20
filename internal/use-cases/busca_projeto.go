package usecases

import (
	"raw-sqlite/internal/database"
	"time"
)

type buscaProjeto struct {
	ProjectRepository database.ProjetoRepository
}

type BuscaProjetoIn struct {
	ID int
}

type BuscaProjetoOut struct {
	ID          int
	Nome        string
	Descricao   string
	DataInicio  time.Time
	DataTermino time.Time
	Status      string
}

func (b *buscaProjeto) Execute(in BuscaProjetoIn) (BuscaProjetoOut, error) {
	projeto, err := b.ProjectRepository.GetProjetoByID(in.ID)
	if err != nil {
		return BuscaProjetoOut{}, err
	}

	return BuscaProjetoOut{
		ID:          projeto.ID,
		Nome:        projeto.Nome,
		Descricao:   projeto.Descricao,
		DataInicio:  projeto.DataInicio,
		DataTermino: projeto.DataTermino,
		Status:      string(projeto.Status),
	}, nil
}
