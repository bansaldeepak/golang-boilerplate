package main

import (
	"golang-boilerplate/config/database"
	"golang-boilerplate/config/environment"
	"golang-boilerplate/config/logger"
	"golang-boilerplate/config/telemetry"
	"golang-boilerplate/docs"
	"golang-boilerplate/modules"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Load configuration
	cfg := environment.LoadConfig()

	// // Set up logger
	// logger, sugar := logger.NewLogger()
	// defer logger.Sync()

	sugar := logger.Sugar

	// Set up OpenTelemetry
	tp, err := telemetry.SetupTelemetry(sugar)
	if err != nil {
		sugar.Fatal(err)
	}
	defer telemetry.ShutdownTelemetry(tp)

	sugar.Infof("Loaded environment: %+v", cfg)

	// Set up database connection
	err = database.Connect(cfg)
	if err != nil {
		sugar.Fatal("failed to connect database")
	}

	database.InitDB(cfg)

	// Set up Gin router
	// r := gin.Default()
	r := modules.GetRouter()

	// Define routes
	// r.GET("/api/v1/health", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"message": "ok"})
	// })
	// r.GET("/api/v1/users", userController.GetUsers)
	// healthController := health.NewHealthController()
	// r.GET("/health/status", healthController.GetStatus)

	docs.SwaggerInfo.BasePath = "/v1"

	// Redirect /swagger to /swagger/index.html
	r.GET("/#docs", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})

	// Serve Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Start server
	sugar.Infof("Starting server on :%s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		sugar.Fatal(err)
	}
}
