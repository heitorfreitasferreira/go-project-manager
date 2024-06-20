package usecases

import "raw-sqlite/internal/database"

var (
	BuscaProjeto       buscaProjeto
	CreateProject      createProject
	BuscaTodosProjetos buscaTodosProjetos
)

func Init(repos database.Service) {
	BuscaProjeto = buscaProjeto{ProjectRepository: repos}
	CreateProject = createProject{ProjectRepository: repos}
	BuscaTodosProjetos = buscaTodosProjetos{ProjectRepository: repos}
}
