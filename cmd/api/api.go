package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/lelandaure/ecom/service/user"
	"log"
	"net/http"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.
		PathPrefix("/api/v1").
		Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subRouter)

	log.Println("Listening on", s.address)

	return http.ListenAndServe(s.address, nil)
}
