package models

import (
	"database/sql"
)

type TaskStatus string

const (
	NotStarted TaskStatus = "NOT_STARTED"
	InProgress TaskStatus = "IN_PROGRESS"
	Completed  TaskStatus = "COMPLETED"
)

type Tarefa struct {
	ID          int
	Nome        sql.NullString
	Descricao   sql.NullString
	Responsavel sql.NullString
	DataInicio  sql.NullTime
	DataTermino sql.NullTime
	Status      TaskStatus
	ProjetoID   int
}
