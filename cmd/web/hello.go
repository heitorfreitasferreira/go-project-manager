package web

import (
	"log"
	"net/http"

	usecases "github.com/heitorfreitasferreira/go-project-manager/internal/use-cases"
)

func HelloWebHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	name := r.FormValue("name")
	component := HelloPost(name)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}

func ViewProjectsHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := usecases.FindAllProjects.Execute(usecases.FindProjectsIn{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	component := ProjectNeehVidaahh(projects)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("Error rendering in ViewProjectsHandler: %e", err)
	}
}
