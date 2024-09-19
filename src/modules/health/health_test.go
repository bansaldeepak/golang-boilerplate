package health

import (
	"golang-boilerplate/test"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	test.TestMain(m) // Call the global TestMain from setup.go
}

func TestGetStatus(t *testing.T) {
	// Set up Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Initialize HealthController
	healthController := NewHealthController()
	router.GET("/api/v1/health", healthController.GetStatus)

	// Create a request to send to the above route
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/health", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "database_status")
	assert.Contains(t, w.Body.String(), "http_status")
}
