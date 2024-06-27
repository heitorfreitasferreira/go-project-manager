package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/heitorfreitasferreira/go-project-manager/internal/models"
)

// ProjectRepository define os métodos para manipular Projetos
type ProjectRepository interface {
	CreateProject(projeto *models.Project) error
	GetProjectoByID(id int) (*models.Project, error)
	UpdateProject(projeto *models.Project) error
	DeleteProject(id int) error
	GetAllProject() ([]*models.Project, error)
}

func (s *service) CreateProject(projeto *models.Project) error {
	_, err := s.db.Exec("INSERT INTO projects (`name`, `description`, `start_date`, `end_date`, `status`) VALUES (?, ?, ?, ?, ?)",
		projeto.Name, projeto.Description, projeto.StartDate, projeto.EndDate, projeto.Status)
	return err
}
func (s *service) GetProjectoByID(id int) (*models.Project, error) {
	query := `SELECT p.id, p.name, p.description, p.start_date, p.end_date, p.status, t.id, t.name, t.description, t.responsible, t.start_date, t.end_date, t.status, t.project_id
			  FROM projects p
			  LEFT JOIN tasks t ON p.id = t.project_id 
			  WHERE p.id = ?`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projeto := &models.Project{}
	tasksMap := make(map[int64]*models.Task)

	for rows.Next() {
		var task models.Task
		err = rows.Scan(&projeto.ID, &projeto.Name, &projeto.Description, &projeto.StartDate, &projeto.EndDate, &projeto.Status,
			&task.ID, &task.Name, &task.Description, &task.Owner, &task.StartDate, &task.EndDate, &task.Status, &task.ProjectId)
		if err != nil {
			log.Printf("error scanning project and task. Err: %v", err)
			return nil, err
		}

		if task.ID.Valid && tasksMap[task.ID.Int64] == nil { // Verifica se a tarefa já foi adicionada para evitar duplicatas
			tasksMap[task.ID.Int64] = &task
		}
	}

	// Converte o map de tarefas para uma slice e atribui ao projeto
	for _, task := range tasksMap {
		projeto.Tasks = append(projeto.Tasks, task)
	}

	if projeto.ID == 0 { // Verifica se um projeto foi encontrado
		return nil, fmt.Errorf("projeto não encontrado")
	}

	return projeto, nil
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
	query := `SELECT p.id, p.name, p.description, p.start_date, p.end_date, p.status FROM projects p`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projetosMap := make(map[int]*models.Project)
	for rows.Next() {
		projeto := &models.Project{}
		err = rows.Scan(&projeto.ID, &projeto.Name, &projeto.Description, &projeto.StartDate, &projeto.EndDate, &projeto.Status)
		if err != nil {
			log.Default().Printf("error scanning project Err: %v", err)
			return nil, err
		}

		projetosMap[projeto.ID] = projeto
	}

	projetos, err := s.loadAllTasks(projetosMap)
	if err != nil {
		return nil, errors.New("error loading tasks")
	}
	return projetos, nil
}

func (s *service) loadAllTasks(projectMap map[int]*models.Project) ([]*models.Project, error) {
	query := `SELECT id, name, description, responsible, start_date, end_date, status, project_id FROM tasks`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.ID, &task.Name, &task.Description, &task.Owner, &task.StartDate, &task.EndDate, &task.Status, &task.ProjectId)
		if err != nil {
			log.Printf("error scanning task. Err: %v", err)
			return nil, err
		}

		if _, ok := projectMap[task.ProjectId]; ok {
			if projectMap[task.ProjectId].Tasks == nil {
				projectMap[task.ProjectId].Tasks = []*models.Task{
					&task,
				}
				continue
			}
			projectMap[task.ProjectId].Tasks = append(projectMap[task.ProjectId].Tasks, &task)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	var projetos []*models.Project
	for _, projeto := range projectMap {
		projetos = append(projetos, projeto)
	}
	return projetos, nil
}
