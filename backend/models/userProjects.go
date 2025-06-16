package models

import "gorm.io/gorm"

type UserProjects struct {
	gorm.Model
	UserID    uint     `gorm:"not null"`
	ProjectID uint     `gorm:"not null"`
	Role      string   `gorm:"type:varchar(50);not null;default:'member'"` // Role in project like 'owner', 'member', etc
	User      Users    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Project   Projects `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
