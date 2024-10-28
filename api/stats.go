package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getStats() gin.HandlerFunc {
	return func(c *gin.Context) {
		stats, err := s.service.GetStats()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, stats)
	}
}
