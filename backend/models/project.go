package models

import "gorm.io/gorm"

type Projects struct {
	gorm.Model

	Name        string `json:"name" gorm:"uniqueIndex"`
	Description string `json:"description"`

	OwnerID int `json:"owner_id" gorm:"index"`
}
