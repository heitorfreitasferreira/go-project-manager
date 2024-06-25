package server

import (
	"encoding/json"
	"net/http"

	"github.com/heitorfreitasferreira/go-project-manager/cmd/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	r.Get("/health", s.healthHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)
	r.Get("/hello", web.BaseWebHandler)
	r.Mount("/api", s.apiRouter())
	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) apiRouter() http.Handler {
	r := chi.NewRouter()

	r.Mount("/project", s.projetoRouter())
	r.Mount("/task", s.taskRouter())
	return r
}
