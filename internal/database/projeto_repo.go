package database

import "github.com/heitorfreitasferreira/go-project-manager/internal/models"

// ProjectRepository define os métodos para manipular Projetos
type ProjectRepository interface {
	CreateProject(projeto *models.Project) error
	GetProjectoByID(id int) (*models.Project, error)
	UpdateProject(projeto *models.Project) error
	DeleteProject(id int) error
	GetAllProject() ([]*models.Project, error)
	LoadTasks(projeto *models.Project) error
}

func (s *service) CreateProject(projeto *models.Project) error {
	_, err := s.db.Exec("INSERT INTO projects (`name`, `description`, `start_date`, `end_date`, `status`) VALUES (?, ?, ?, ?, ?)",
		projeto.Name, projeto.Description, projeto.StartDate, projeto.EndDate, projeto.Status)
	return err
}

func (s *service) GetProjectoByID(id int) (*models.Project, error) {
	projeto := &models.Project{}
	err := s.db.QueryRow("SELECT * FROM projects WHERE id = ?", id).Scan(&projeto.ID, &projeto.Name, &projeto.Description, &projeto.StartDate, &projeto.EndDate, &projeto.Status)
	if err != nil {
		return nil, err
	}
	return projeto, nil
}
func (s *service) LoadTasks(projeto *models.Project) error {
	rows, err := s.db.Query("SELECT * FROM tasks WHERE project_id = ?", projeto.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		t := &models.Task{}
		err = rows.Scan(&t.ID, &t.Name, &t.Description, &t.StartDate, &t.EndDate, &t.Status, &t.ProjectId)
		if err != nil {
			return err
		}
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	projeto.Tasks = tasks
	return nil
}

func (s *service) UpdateProject(projeto *models.Project) error {
	_, err := s.db.Exec("UPDATE projects SET `name` = ?, `description` = ?, `start_date` = ?, `end_date` = ?, `status` = ? WHERE id = ?",
		projeto.Name, projeto.Description, projeto.StartDate, projeto.EndDate, projeto.Status, projeto.ID)
	return err
}

func (s *service) DeleteProject(id int) error {
	_, err := s.db.Exec("DELETE FROM projects WHERE id = ?", id)
	return err
}

func (s *service) GetAllProject() ([]*models.Project, error) {
	rows, err := s.db.Query("SELECT * FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projetos []*models.Project
	for rows.Next() {
		projeto := &models.Project{}
		err = rows.Scan(&projeto.ID, &projeto.Name, &projeto.Description, &projeto.StartDate, &projeto.EndDate, &projeto.Status)
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
