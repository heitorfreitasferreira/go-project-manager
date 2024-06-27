package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/heitorfreitasferreira/go-project-manager/internal/types"
	usecases "github.com/heitorfreitasferreira/go-project-manager/internal/use-cases"
	"github.com/heitorfreitasferreira/go-project-manager/internal/views"
)

func (s *Server) htmxProjectRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", s.listAllProjectsHtmxHandler)
	r.Post("/", s.createProjectHtmxHandler)
	return r
}

func (s *Server) createProjectHtmxHandler(w http.ResponseWriter, r *http.Request) {
	// Parseia o formulário recebido
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	startDate, err := time.Parse("2006-01-02", r.FormValue("start_date"))

	if err != nil {
		http.Error(w, "Invalid start date", http.StatusBadRequest)

		return
	}
	endDate, err := time.Parse("2006-01-02", r.FormValue("end_date"))
	if err != nil {
		http.Error(w, "Invalid end date", http.StatusBadRequest)
		return
	}
	in := usecases.CreateProjectIn{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		StartDate:   startDate,
		EndDate:     endDate,
		Status:      r.FormValue("status"),
	}

	out, err := usecases.CreateProject.Execute(in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := views.ViewSingleProject(types.Project(out))

	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error rendering in createProjectHtmxHandler: %e", err)
		return
	}
}

func (s *Server) listAllProjectsHtmxHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := usecases.FindAllProjects.Execute(usecases.FindProjectsIn{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	component := views.ViewAllProjects(projects)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error rendering in ViewProjectsHandler: %e", err)
	}
}
