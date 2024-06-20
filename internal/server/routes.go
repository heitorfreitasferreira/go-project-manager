package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"fmt"
	"time"

	"raw-sqlite/cmd/web"
	usecases "raw-sqlite/internal/use-cases"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"nhooyr.io/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	r.Get("/websocket", s.websocketHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/assets/*", fileServer)
	r.Get("/web", templ.Handler(web.HelloForm()).ServeHTTP)
	r.Post("/hello", web.HelloWebHandler)

	r.Post("/project", s.createProjectHandler)
	r.Get("/project/{id}", s.getProjectsHandler)
	r.Get("/projects", s.getAllProjectsHandler)
	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
		err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
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
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error converting id to int. Err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
		return
	}

	project, err := usecases.BuscaProjeto.Execute(usecases.BuscaProjetoIn{ID: idInt})
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
