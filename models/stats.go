package models

import "time"

type Stats struct {
	ID        string    `json:"id"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"created_at"`
}
