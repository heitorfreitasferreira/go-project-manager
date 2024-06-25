package usecases

import (
	"raw-sqlite/internal/database"
	"time"
)

type findProjectByID struct {
	projectRepository database.ProjectRepository
}

type FindProjectByIdIn int

type FindProjectByIdOut struct {
	ID          int
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Status      string
}

func (b *findProjectByID) Execute(in FindProjectByIdIn) (FindProjectByIdOut, error) {
	projeto, err := b.projectRepository.GetProjectoByID(int(in))
	if err != nil {
		return FindProjectByIdOut{}, err
	}

	return FindProjectByIdOut{
		ID:          projeto.ID,
		Name:        projeto.Name.String,
		Description: projeto.Description.String,
		StartDate:   projeto.StartDate.Time,
		EndDate:     projeto.EndDate.Time,
		Status:      string(projeto.Status),
	}, nil
}
