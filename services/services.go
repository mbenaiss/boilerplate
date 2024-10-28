package services

import "github.com/mbenaiss/boilerplate-go/db"

// Service is the  service
type Service struct {
	db db.DB
}

// NewService creates a new stats service
func NewService(db db.DB) *Service {
	return &Service{db: db}
}
