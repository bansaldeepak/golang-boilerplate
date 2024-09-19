package health

import (
	"golang-boilerplate/config/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	Service *HealthService
}

func NewHealthController() *HealthController {
	db := database.GetDB()
	healthService := &HealthService{DB: db}
	return &HealthController{Service: healthService}
}

// GetStatus godoc
// @Summary Get the health status of the service
// @Description Get the health status of the service
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /health/status [get]
func (ctrl *HealthController) GetStatus(c *gin.Context) {
	status, err := ctrl.Service.GetStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"database_status": status, "http_status": http.StatusOK},
	)
}
