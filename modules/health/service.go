package health

import (
	"gorm.io/gorm"
)

type HealthService struct {
	DB *gorm.DB
}

func (s *HealthService) GetStatus() (string, error) {
	sqlDB, err := s.DB.DB()
	if err != nil {
		return "", err
	}
	err = sqlDB.Ping()
	if err != nil {
		return "", err
	}
	return "ok", nil
}
