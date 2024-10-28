package db

import "github.com/mbenaiss/boilerplate-go/models"

type DB interface {
	GetStats() (*models.Stats, error)
}
