package usecases

import "raw-sqlite/internal/database"

var (
	BuscaProjeto       buscaProjeto
	CreateProject      createProject
	BuscaTodosProjetos buscaTodosProjetos

	FindAllTasks findAllTasks
	CreateTask   createTask
)

func Init(repos database.Service) {
	BuscaProjeto = buscaProjeto{projectRepository: repos}
	CreateProject = createProject{projectRepository: repos}
	BuscaTodosProjetos = buscaTodosProjetos{projectRepository: repos}
	FindAllTasks = findAllTasks{taskRepository: repos}
	CreateTask = createTask{taskRepository: repos}
}
