package models

import "time"

type TarefaStatus string

const (
	NotStarted TarefaStatus = "NOT_STARTED"
	InProgress TarefaStatus = "IN_PROGRESS"
	Completed  TarefaStatus = "COMPLETED"
)

type Tarefa struct {
	ID          int
	Nome        string
	Descricao   string
	Responsavel string
	DataInicio  time.Time
	DataTermino time.Time
	Status      TarefaStatus
	ProjetoID   int
}
