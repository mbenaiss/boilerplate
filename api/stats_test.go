package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mbenaiss/boilerplate-go/models"
)

type MockService struct {
	mock.Mock
}

func (s *MockService) GetStats() (*models.Stats, error) {
	args := s.Called()
	return args.Get(0).(*models.Stats), args.Error(1)
}

func TestGetStats(t *testing.T) {
	service := &MockService{}
	service.On("GetStats").Return(&models.Stats{
		ID:        "1",
		Count:     10,
		CreatedAt: time.Now(),
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	w := httptest.NewRecorder()

	server := NewServer(service)
	router := gin.Default()
	router.GET("/stats", server.getStats())
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
