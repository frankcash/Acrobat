package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

func (s *Server) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s passwod=%s dbname=%s", user, password, dbname)

	var err error
	s.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database connection established")
	}

	s.Router = mux.NewRouter()
	s.InitializeRoutes()
}

func (s *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(":80", s.Router))
}

func (s *Server) InitializeRoutes() {
	s.Router.HandleFunc("/api/v1/items/all", s.allItems).Methods("GET")
	s.Router.HandleFunc("/api/v1/items/{id:[0-9]+}", s.getItemByID).Methods("GET")
	s.Router.HandleFunc("/api/v1/items/new", s.newEntry).Methods("POST")
}
