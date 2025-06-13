package models

import "gorm.io/gorm"

type Boards struct {
	gorm.Model

	Name      string `json:"name" gorm:"unique"`
	Color     string `json:"color" gorm:"unique"`
	ProjectID int    `json:"project_id" gorm:"unique"`
}
