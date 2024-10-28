package sqlite

import "github.com/mbenaiss/boilerplate-go/models"

// Client is the sqlite client
type Client struct {
	path string
}

// NewClient creates a new sqlite client
func NewClient(path string) *Client {
	return &Client{path: path}
}

// GetStats gets the stats
func (c *Client) GetStats() (*models.Stats, error) {
	return nil, nil
}
