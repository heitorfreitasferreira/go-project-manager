package server

import (
	"encoding/json"
	"net/http"
	usecases "github.com/heitorfreitasferreira/go-project-manager/internal/use-cases"

	"github.com/go-chi/chi/v5"
)

func (s *Server) taskRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.listTasks)
	r.Post("/", s.createTask)
	// r.Put("/{id}", s.updateTask)
	// r.Delete("/{id}", s.deleteTask)

	return r
}

func (s *Server) listTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := usecases.FindAllTasks.Execute(usecases.FindAllTasksIn{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)

	}
	jsonResponse, _ := json.Marshal(tasks)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (s *Server) createTask(w http.ResponseWriter, r *http.Request) {
	var task usecases.CreateTaskIn
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"error": err.Error()}
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)
		return
	}

	_, err = usecases.CreateTask.Execute(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": err.Error()}
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
