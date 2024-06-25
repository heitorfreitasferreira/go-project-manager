package server

import (
	"encoding/json"
	"log"
	"net/http"
	usecases "raw-sqlite/internal/use-cases"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (s *Server) projetoRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", s.createProjectHandler)
	r.Get("/{id}", s.getProjectsHandler)
	r.Get("/", s.getAllProjectsHandler)
	return r
}

func (s *Server) createProjectHandler(w http.ResponseWriter, r *http.Request) {
	var project usecases.CreateProjectIn
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		log.Printf("error decoding project. Err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	out, err := usecases.CreateProject.Execute(project)
	log.Printf("project created. ID: %d", out.ID)
	if err != nil {
		log.Printf("error creating project. Err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) getProjectsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("error converting id to int. Err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
		return
	}

	project, err := usecases.BuscaProjeto.Execute(usecases.BuscaProjetoIn(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonResp, _ := json.Marshal(project)
	_, _ = w.Write(jsonResp)
}

func (s *Server) getAllProjectsHandler(w http.ResponseWriter, _ *http.Request) {
	projects, err := usecases.BuscaTodosProjetos.Execute(usecases.BuscaTodosProjetosIn{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonResp, _ := json.Marshal(projects)
	_, _ = w.Write(jsonResp)
}
