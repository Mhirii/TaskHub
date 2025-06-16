package infra

import (
	"time"

	"github.com/Mhirii/TaskHub/backend/models"
	"github.com/Mhirii/TaskHub/backend/pkg/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAdminUser(db *gorm.DB, cfg config.Config) error {
	admincfg := cfg.Seed.Admin
	pw, err := bcrypt.GenerateFromPassword([]byte(admincfg.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	adminUser := models.Users{
		Model: gorm.Model{
			ID:        uint(1),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: admincfg.Username,
		Email:    admincfg.Email,
		Password: string(pw),
		Roles:    "admin",
	}

	var existingUser models.Users
	result := db.Where("username = ?", "admin").First(&existingUser)
	if result.Error == nil {
		return nil
	}
	if result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}

	return db.Create(&adminUser).Error
}
