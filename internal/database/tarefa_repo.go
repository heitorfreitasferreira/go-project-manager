package database

import "raw-sqlite/internal/models"

type TaskRepository interface {
	CreateTask(tarefa *models.Tarefa) error
	GetTaskByID(id int) (*models.Tarefa, error)
	UpdateTask(tarefa *models.Tarefa) error
	DeleteTask(id int) error
	GetAllTasksByProjectID(projetoID int) ([]*models.Tarefa, error)
	GetAllTasks() ([]*models.Tarefa, error)
}

func (s *service) CreateTask(tarefa *models.Tarefa) error {
	_, err := s.db.Exec("INSERT INTO tasks (name, description, start_date, end_date, status, project_id) VALUES (?, ?, ?, ?, ?, ?)",
		tarefa.Nome, tarefa.Descricao, tarefa.DataInicio, tarefa.DataTermino, tarefa.Status, tarefa.ProjetoID)
	return err
}

func (s *service) GetTaskByID(id int) (*models.Tarefa, error) {
	tarefa := &models.Tarefa{}
	err := s.db.QueryRow("SELECT * FROM tasks WHERE id = ?", id).Scan(&tarefa.ID, &tarefa.Nome, &tarefa.Descricao, &tarefa.DataInicio, &tarefa.DataTermino, &tarefa.Status, &tarefa.ProjetoID)
	if err != nil {
		return nil, err
	}
	return tarefa, nil
}

func (s *service) UpdateTask(tarefa *models.Tarefa) error {
	_, err := s.db.Exec("UPDATE tasks SET name = ?, description = ?, start_date = ?, end_date = ?, status = ?, project_id = ? WHERE id = ?",
		tarefa.Nome, tarefa.Descricao, tarefa.DataInicio, tarefa.DataTermino, tarefa.Status, tarefa.ProjetoID, tarefa.ID)
	return err
}

func (s *service) DeleteTask(id int) error {
	_, err := s.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

func (s *service) GetAllTasksByProjectID(projetoID int) ([]*models.Tarefa, error) {
	rows, err := s.db.Query("SELECT * FROM tasks WHERE project_id = ?", projetoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Tarefa
	for rows.Next() {
		tarefa := &models.Tarefa{}
		err = rows.Scan(&tarefa.ID, &tarefa.Nome, &tarefa.Descricao, &tarefa.DataInicio, &tarefa.DataTermino, &tarefa.Status, &tarefa.ProjetoID)
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

func (s *service) GetAllTasks() ([]*models.Tarefa, error) {
	rows, err := s.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Tarefa
	for rows.Next() {
		tarefa := &models.Tarefa{}
		err = rows.Scan(&tarefa.ID, &tarefa.Nome, &tarefa.Descricao, &tarefa.Responsavel, &tarefa.DataInicio, &tarefa.DataTermino, &tarefa.Status, &tarefa.ProjetoID)
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
