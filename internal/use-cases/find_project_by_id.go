package usecases

import (
	"raw-sqlite/internal/database"
	"time"
)

type buscaProjeto struct {
	projectRepository database.ProjectRepository
}

type BuscaProjetoIn int

type BuscaProjetoOut struct {
	ID          int
	Nome        string
	Descricao   string
	DataInicio  time.Time
	DataTermino time.Time
	Status      string
}

func (b *buscaProjeto) Execute(in BuscaProjetoIn) (BuscaProjetoOut, error) {
	projeto, err := b.projectRepository.GetProjectoByID(int(in))
	if err != nil {
		return BuscaProjetoOut{}, err
	}

	return BuscaProjetoOut{
		ID:          projeto.ID,
		Nome:        projeto.Name.String,
		Descricao:   projeto.Description.String,
		DataInicio:  projeto.StartDate.Time,
		DataTermino: projeto.EndDate.Time,
		Status:      string(projeto.Status),
	}, nil
}
