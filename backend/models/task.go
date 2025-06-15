package models

import "gorm.io/gorm"

type Tasks struct {
	gorm.Model

	Name        string `json:"name"`
	Description string `json:"description"`

	BoardID   int `json:"board_id"`
	ProjectID int `json:"project_id"`

	Order      int    `json:"order"`
	ParentID   int    `json:"parent_id"`
	Status     int    `json:"status" gorm:"default:0"`   // 0 -> open, 1 -> closed
	Priority   *int   `json:"priority" gorm:"default:0"` // 0 -> low, 1 -> medium, 2 -> high
	DueDate    *int   `json:"due_date"`                  // unix timestamp
	AssigneeID *int   `json:"assignee"`                  // 0 -> unassigned
	Tags       string `json:"tags" gorm:"type:text"`

	CreatedBy int `json:"created_by"`
}
