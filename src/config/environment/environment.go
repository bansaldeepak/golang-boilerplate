package environment

import (
	"log"
	"os"
)

type Config struct {
	DBHost      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBPort      string
	AppPort     string
	Environment string
}

func LoadConfig() Config {
	return Config{
		DBHost:      getEnv("DB_HOST", ""),
		DBUser:      getEnv("DB_USER", ""),
		DBPassword:  getEnv("DB_PASSWORD", ""),
		DBName:      getEnv("DB_NAME", ""),
		DBPort:      getEnv("DB_PORT", ""),
		AppPort:     getEnv("APP_PORT", ""),
		Environment: getEnv("ENVIRONMENT", ""),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		if defaultValue == "" {
			log.Fatalf("Environment variable %s not set", key)
		}
		return defaultValue
	}
	return value
}
