package models

import "time"

type ProjectStatus string

const (
	ProjectPlanning   ProjectStatus = "PLANNING"
	ProjectInProgress ProjectStatus = "IN_PROGRESS"
	ProjectDone       ProjectStatus = "DONE"
)

type Projeto struct {
	ID          int
	Nome        string
	Descricao   string
	DataInicio  time.Time
	DataTermino time.Time
	Status      ProjectStatus
}

func ToProjectStatus(status string) ProjectStatus {
	switch status {
	case string(ProjectPlanning):
		return ProjectPlanning
	case string(ProjectInProgress):
		return ProjectInProgress
	case string(ProjectDone):
		return ProjectDone
	default:
		return ProjectPlanning
	}
}
