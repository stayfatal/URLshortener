package server

import "url/internal/database"

type Server struct {
	DBManager *database.DBManager
	domain    string
}

func NewServer() *Server {
	return &Server{DBManager: database.NewDBManager(), domain: "http://localhost:8080/"}
}
