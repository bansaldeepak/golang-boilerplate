package test

import (
	"log"
	"os"
	"testing"

	"golang-boilerplate/config/database"
	"golang-boilerplate/config/environment"
)

func GlobalSetup() {
	// Load test environment variables
	cfg := environment.LoadConfig()

	// Connect to the test database
	err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	// Initialize the database with test models
	database.InitDB(cfg)
}

func TestMain(m *testing.M) {
	GlobalSetup()
	code := m.Run()
	os.Exit(code)
}
