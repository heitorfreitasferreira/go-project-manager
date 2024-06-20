package database

import "raw-sqlite/internal/models"

type TarefaRepository interface {
	CreateTarefa(tarefa *models.Tarefa) error
	GetTarefaByID(id int) (*models.Tarefa, error)
	UpdateTarefa(tarefa *models.Tarefa) error
	DeleteTarefa(id int) error
	GetAllTarefaByProjetoID(projetoID int) ([]*models.Tarefa, error)
	GetAllTarefa() ([]*models.Tarefa, error)
}

func (s *service) CreateTarefa(tarefa *models.Tarefa) error {
	_, err := s.db.Exec("INSERT INTO tarefas (nome, descricao, data_inicio, data_termino, status, projeto_id) VALUES (?, ?, ?, ?, ?, ?)",
		tarefa.Nome, tarefa.Descricao, tarefa.DataInicio, tarefa.DataTermino, tarefa.Status, tarefa.ProjetoID)
	return err
}

func (s *service) GetTarefaByID(id int) (*models.Tarefa, error) {
	tarefa := &models.Tarefa{}
	err := s.db.QueryRow("SELECT * FROM tarefas WHERE id = ?", id).Scan(&tarefa.ID, &tarefa.Nome, &tarefa.Descricao, &tarefa.DataInicio, &tarefa.DataTermino, &tarefa.Status, &tarefa.ProjetoID)
	if err != nil {
		return nil, err
	}
	return tarefa, nil
}

func (s *service) UpdateTarefa(tarefa *models.Tarefa) error {
	_, err := s.db.Exec("UPDATE tarefas SET nome = ?, descricao = ?, data_inicio = ?, data_termino = ?, status = ?, projeto_id = ? WHERE id = ?",
		tarefa.Nome, tarefa.Descricao, tarefa.DataInicio, tarefa.DataTermino, tarefa.Status, tarefa.ProjetoID, tarefa.ID)
	return err
}

func (s *service) DeleteTarefa(id int) error {
	_, err := s.db.Exec("DELETE FROM tarefas WHERE id = ?", id)
	return err
}

func (s *service) GetAllTarefaByProjetoID(projetoID int) ([]*models.Tarefa, error) {
	rows, err := s.db.Query("SELECT * FROM tarefas WHERE projeto_id = ?", projetoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tarefas []*models.Tarefa
	for rows.Next() {
		tarefa := &models.Tarefa{}
		err = rows.Scan(&tarefa.ID, &tarefa.Nome, &tarefa.Descricao, &tarefa.DataInicio, &tarefa.DataTermino, &tarefa.Status, &tarefa.ProjetoID)
		if err != nil {
			return nil, err
		}
		tarefas = append(tarefas, tarefa)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tarefas, nil
}

func (s *service) GetAllTarefa() ([]*models.Tarefa, error) {
	rows, err := s.db.Query("SELECT * FROM tarefas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tarefas []*models.Tarefa
	for rows.Next() {
		tarefa := &models.Tarefa{}
		err = rows.Scan(&tarefa.ID, &tarefa.Nome, &tarefa.Descricao, &tarefa.DataInicio, &tarefa.DataTermino, &tarefa.Status, &tarefa.ProjetoID)
		if err != nil {
			return nil, err
		}
		tarefas = append(tarefas, tarefa)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tarefas, nil
}
