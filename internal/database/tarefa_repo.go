package database

import "raw-sqlite/internal/models"

type TaskRepository interface {
	CreateTask(tarefa *models.Task) error
	GetTaskByID(id int) (*models.Task, error)
	UpdateTask(tarefa *models.Task) error
	DeleteTask(id int) error
	GetAllTasksByProjectID(projetoID int) ([]*models.Task, error)
	GetAllTasks() ([]*models.Task, error)
}

func (s *service) CreateTask(tarefa *models.Task) error {
	_, err := s.db.Exec("INSERT INTO tasks (name, description, start_date, end_date, status, project_id) VALUES (?, ?, ?, ?, ?, ?)",
		tarefa.Name, tarefa.Description, tarefa.StartDate, tarefa.EndDate, tarefa.Status, tarefa.ProjectId)
	return err
}

func (s *service) GetTaskByID(id int) (*models.Task, error) {
	tarefa := &models.Task{}
	err := s.db.QueryRow("SELECT * FROM tasks WHERE id = ?", id).Scan(&tarefa.ID, &tarefa.Name, &tarefa.Description, &tarefa.StartDate, &tarefa.EndDate, &tarefa.Status, &tarefa.ProjectId)
	if err != nil {
		return nil, err
	}
	return tarefa, nil
}

func (s *service) UpdateTask(tarefa *models.Task) error {
	_, err := s.db.Exec("UPDATE tasks SET name = ?, description = ?, start_date = ?, end_date = ?, status = ?, project_id = ? WHERE id = ?",
		tarefa.Name, tarefa.Description, tarefa.StartDate, tarefa.EndDate, tarefa.Status, tarefa.ProjectId, tarefa.ID)
	return err
}

func (s *service) DeleteTask(id int) error {
	_, err := s.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

func (s *service) GetAllTasksByProjectID(projetoID int) ([]*models.Task, error) {
	rows, err := s.db.Query("SELECT * FROM tasks WHERE project_id = ?", projetoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		tarefa := &models.Task{}
		err = rows.Scan(&tarefa.ID, &tarefa.Name, &tarefa.Description, &tarefa.StartDate, &tarefa.EndDate, &tarefa.Status, &tarefa.ProjectId)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, tarefa)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *service) GetAllTasks() ([]*models.Task, error) {
	rows, err := s.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		tarefa := &models.Task{}
		err = rows.Scan(&tarefa.ID, &tarefa.Name, &tarefa.Description, &tarefa.Owner, &tarefa.StartDate, &tarefa.EndDate, &tarefa.Status, &tarefa.ProjectId)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, tarefa)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
