package api

import (
	"database/sql"
	"net/http"

	"github.com/kannape/go-e-commerce-app/service/user"
)

type APIServer struct {
	addr     string
	database *sql.DB
}

func NewAPIServer(addr string, database *sql.DB) *APIServer {
	return &APIServer{
		addr:     addr,
		database: database,
	}
}

func (s *APIServer) Run() error {
	router := http.DefaultServeMux
	subrouter := http.DefaultServeMux

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	router.Handle("/api/v1/", http.StripPrefix("/api/v1", subrouter))

	return http.ListenAndServe(s.addr, router)
}
