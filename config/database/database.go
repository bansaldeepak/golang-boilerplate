package database

import (
	"golang-boilerplate/config/environment"
	"golang-boilerplate/config/logger"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

var sugar = logger.Sugar

func Connect(cfg environment.Config) error {
	var err error
	once.Do(func() {
		dsn := "host=" + cfg.DBHost + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " port=" + cfg.DBPort + " sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			sugar.Errorf("Failed to connect to database: %v", err)
		}
	})
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func InitDB(cfg environment.Config, models ...interface{}) {
	if cfg.Environment == "development" {
		sugar.Info("Initializing database...")
		err := db.AutoMigrate(models...)
		if err != nil {
			sugar.Fatalf("Failed to migrate database: %v", err)
		}
		sugar.Info("Database initialization complete.")
	} else {
		sugar.Info("Database initialization skipped.")
	}
}
