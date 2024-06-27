package web

import (
	"log"
	"net/http"

	usecases "github.com/heitorfreitasferreira/go-project-manager/internal/use-cases"
	"github.com/heitorfreitasferreira/go-project-manager/internal/views"
)

func ViewProjectsHandler(w http.ResponseWriter, r *http.Request) {
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
