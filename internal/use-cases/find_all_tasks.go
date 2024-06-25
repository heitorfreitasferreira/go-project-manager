package usecases

import (
	"github.com/heitorfreitasferreira/go-project-manager/internal/database"
	"github.com/heitorfreitasferreira/go-project-manager/internal/types"
)

type findAllTasks struct {
	taskRepository database.TaskRepository
}
type FindAllTasksIn struct {
}
type FindAllTasksOut []types.Task

func (u findAllTasks) Execute(in FindAllTasksIn) (FindAllTasksOut, error) {
	tasks, err := u.taskRepository.GetAllTasks()
	if err != nil {
		return nil, err
	}
	tasksDto := make(FindAllTasksOut, len(tasks))

	for i, task := range tasks {
		tasksDto[i] = types.FromModelToTask(*task)
	}
	return tasksDto, nil
}
