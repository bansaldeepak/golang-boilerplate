package modules

import (
	"golang-boilerplate/modules/health"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	healthController := health.NewHealthController()

	// Set up Gin router
	r := gin.Default()

	// Define routes
	// r.GET("/health/status", healthController.GetStatus)

	v1 := r.Group("/v1")
	{
		health := v1.Group("/health")
		{
			health.GET("/status", healthController.GetStatus)
		}
	}

	return r
}
