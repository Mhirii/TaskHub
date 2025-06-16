package infra

import (
	"fmt"

	"github.com/Mhirii/TaskHub/backend/models"
	"github.com/Mhirii/TaskHub/backend/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg config.Config) (*gorm.DB, error) {
	DSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.SSLMode)
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Projects{})
	db.AutoMigrate(&models.Boards{})
	db.AutoMigrate(&models.Tasks{})
	db.AutoMigrate(&models.UserProjects{})

	if err := db.Raw("SELECT 1").Error; err != nil {
		return nil, fmt.Errorf("database connection test failed: %w", err)
	}

	return db, nil
}
