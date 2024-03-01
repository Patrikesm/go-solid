package server

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Server struct {
	db *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{db: db}
}

func (s *Server) Run() {
	r := gin.Default()

	if err := s.MapHandlers(r); err != nil {
		return
	}

	r.Run(":3000")
}
