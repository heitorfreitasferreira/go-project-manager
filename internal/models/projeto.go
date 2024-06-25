package models

import (
	"database/sql"
)

type ProjectStatus string

const (
	ProjectPlanning   ProjectStatus = "PLANNING"
	ProjectInProgress ProjectStatus = "IN_PROGRESS"
	ProjectDone       ProjectStatus = "DONE"
)

type Project struct {
	ID          int
	Nome        sql.NullString
	Descricao   sql.NullString
	DataInicio  sql.NullTime
	DataTermino sql.NullTime
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
