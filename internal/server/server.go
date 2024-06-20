package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"raw-sqlite/internal/database"
	usecases "raw-sqlite/internal/use-cases"
)

type Server struct {
	port int

	db database.Service

	createProject   usecases.CreateProject
	findProjectById usecases.BuscaProjeto
	findAllProjects usecases.BuscaTodosProjetos
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := database.New()

	NewServer := &Server{
		port: port,
		db:   db,

		createProject: usecases.CreateProject{
			ProjectRepository: db,
		},
		findProjectById: usecases.BuscaProjeto{
			ProjectRepository: db,
		},
		findAllProjects: usecases.BuscaTodosProjetos{
			ProjectRepository: db,
		},
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
