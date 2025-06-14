package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
	Roles    string `json:"roles"`
}
