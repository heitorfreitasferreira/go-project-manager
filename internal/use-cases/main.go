package usecases

import "github.com/heitorfreitasferreira/go-project-manager/internal/database"

var (
	FindProjectById findProjectByID
	CreateProject   createProject
	FindAllProjects findAllProjects

	FindAllTasks findAllTasks
	CreateTask   createTask
)

func Init(repos database.Service) {
	FindProjectById = findProjectByID{projectRepository: repos}
	CreateProject = createProject{projectRepository: repos}
	FindAllProjects = findAllProjects{projectRepository: repos}
	FindAllTasks = findAllTasks{taskRepository: repos}
	CreateTask = createTask{taskRepository: repos}
}
