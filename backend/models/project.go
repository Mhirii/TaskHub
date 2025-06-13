package models

import "gorm.io/gorm"

type Projects struct {
	gorm.Model

	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description" gorm:"unique"`

	OwnerID int `json:"owner_id" gorm:"unique"`
}
