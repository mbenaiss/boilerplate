package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mbenaiss/boilerplate-go/models"
)

type services interface {
	GetStats() (*models.Stats, error)
}

type Server struct {
	service services
	router  *gin.Engine
}

// NewServer creates a new server
func NewServer(service services) *Server {
	return &Server{service: service}
}

// Start starts the server
func (s *Server) Start() error {
	s.router = gin.Default()
	s.registerRoutes()
	return nil
}

// Stop stops the server
func (s *Server) Stop() error {
	return nil
}

func (s *Server) registerRoutes() {
	s.router.GET("/stats", s.getStats())
}
