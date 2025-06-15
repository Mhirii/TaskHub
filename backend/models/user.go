package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"password"`
	Roles    string `json:"roles"`
}
