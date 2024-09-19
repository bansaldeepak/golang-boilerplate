package environment

import (
	"log"
	"os"
)

type Config struct {
	DBHost         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBPort         string
	AppPort        string
	AppEnvironment string
	AppName        string
}

func LoadConfig() Config {
	return Config{
		DBHost:         getEnv("POSTGRES_HOST", ""),
		DBUser:         getEnv("POSTGRES_USER", ""),
		DBPassword:     getEnv("POSTGRES_PASSWORD", ""),
		DBName:         getEnv("POSTGRES_NAME", ""),
		DBPort:         getEnv("POSTGRES_PORT", ""),
		AppPort:        getEnv("APP_PORT", ""),
		AppEnvironment: getEnv("APP_ENVIRONMENT", ""),
		AppName:        getEnv("APP_NAME", ""),
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
