package database

import "raw-sqlite/internal/models"

// ProjetoRepository define os métodos para manipular Projetos
type ProjetoRepository interface {
	CreateProjeto(projeto *models.Projeto) error
	GetProjetoByID(id int) (*models.Projeto, error)
	UpdateProjeto(projeto *models.Projeto) error
	DeleteProjeto(id int) error
	GetAllProjeto() ([]*models.Projeto, error)
}

func (s *service) CreateProjeto(projeto *models.Projeto) error {
	_, err := s.db.Exec("INSERT INTO projects (`name`, `description`, `start_date`, `end_date`, `status`) VALUES (?, ?, ?, ?, ?)",
		projeto.Nome, projeto.Descricao, projeto.DataInicio, projeto.DataTermino, projeto.Status)
	return err
}

func (s *service) GetProjetoByID(id int) (*models.Projeto, error) {
	projeto := &models.Projeto{}
	err := s.db.QueryRow("SELECT * FROM projects WHERE id = ?", id).Scan(&projeto.ID, &projeto.Nome, &projeto.Descricao, &projeto.DataInicio, &projeto.DataTermino, &projeto.Status)
	if err != nil {
		return nil, err
	}
	return projeto, nil
}

func (s *service) UpdateProjeto(projeto *models.Projeto) error {
	_, err := s.db.Exec("UPDATE projects SET `name` = ?, `description` = ?, `start_date` = ?, `end_date` = ?, `status` = ? WHERE id = ?",
		projeto.Nome, projeto.Descricao, projeto.DataInicio, projeto.DataTermino, projeto.Status, projeto.ID)
	return err
}

func (s *service) DeleteProjeto(id int) error {
	_, err := s.db.Exec("DELETE FROM projects WHERE id = ?", id)
	return err
}

func (s *service) GetAllProjeto() ([]*models.Projeto, error) {
	rows, err := s.db.Query("SELECT * FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projetos []*models.Projeto
	for rows.Next() {
		projeto := &models.Projeto{}
		err = rows.Scan(&projeto.ID, &projeto.Nome, &projeto.Descricao, &projeto.DataInicio, &projeto.DataTermino, &projeto.Status)
		if err != nil {
			return nil, err
		}
		projetos = append(projetos, projeto)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projetos, nil
}
