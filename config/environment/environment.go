package environment

import (
	"log"

	"github.com/spf13/viper"
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
	// Set the path to the directory containing the .env file
	viper.AddConfigPath(".")
	// Set the name of the file (without extension)
	viper.SetConfigName(".env")
	// Set the type of the file
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// var config Config
	// err = viper.Unmarshal(&config)
	// if err != nil {
	// 	log.Fatalf("Unable to decode into struct, %v", err)
	// }

	// return config

	return Config{
		DBHost:      viper.GetString("DB_HOST"),
		DBUser:      viper.GetString("DB_USER"),
		DBPassword:  viper.GetString("DB_PASSWORD"),
		DBName:      viper.GetString("DB_NAME"),
		DBPort:      viper.GetString("DB_PORT"),
		AppPort:     viper.GetString("APP_PORT"),
		Environment: viper.GetString("ENVIRONMENT"),
	}
}
