package dto

import (
	"github.com/Mhirii/TaskHub/backend/models"
)

type CreateTaskRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	BoardID     int      `json:"board_id"`
	ProjectID   int      `json:"project_id"`
	Order       int      `json:"order"`
	ParentID    *int     `json:"parent_id"`
	Status      int      `json:"status"`
	Priority    *int     `json:"priority"`
	DueDate     *int     `json:"due_date"`
	AssigneeID  *int     `json:"assignee"`
	Tags        []string `json:"tags"`
}

func (r *CreateTaskRequest) ToModel() *models.Tasks {
	task := &models.Tasks{
		Name:        r.Name,
		Description: r.Description,
		BoardID:     r.BoardID,
		ProjectID:   r.ProjectID,
		Order:       r.Order,
		Status:      r.Status,
		Priority:    r.Priority,
		DueDate:     r.DueDate,
		AssigneeID:  r.AssigneeID,
		Tags:        r.Tags,
	}
	if r.ParentID != nil {
		task.ParentID = *r.ParentID
	}
	return task
}

type CreateTaskResponse struct {
	ID uint `json:"id"`
}

type UpdateTaskRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	BoardID     *int     `json:"board_id"`
	ProjectID   *int     `json:"project_id"`
	Order       *int     `json:"order"`
	ParentID    *int     `json:"parent_id"`
	Status      *int     `json:"status"`
	Priority    *int     `json:"priority"`
	DueDate     *int     `json:"due_date"`
	AssigneeID  *int     `json:"assignee"`
	Tags        []string `json:"tags"`
}

func (r *UpdateTaskRequest) ToModel() *models.Tasks {
	task := &models.Tasks{}
	if r.Priority != nil {
		task.Priority = r.Priority
	}
	if r.DueDate != nil {
		task.DueDate = r.DueDate
	}
	if r.AssigneeID != nil {
		task.AssigneeID = r.AssigneeID
	}
	if r.Tags != nil {
		task.Tags = r.Tags
	}
	if r.Name != nil {
		task.Name = *r.Name
	}
	if r.Description != nil {
		task.Description = *r.Description
	}
	if r.BoardID != nil {
		task.BoardID = *r.BoardID
	}
	if r.ProjectID != nil {
		task.ProjectID = *r.ProjectID
	}
	if r.Order != nil {
		task.Order = *r.Order
	}
	if r.Status != nil {
		task.Status = *r.Status
	}
	if r.ParentID != nil {
		task.ParentID = *r.ParentID
	}
	return task
}

type UpdateTaskResponse struct {
	ID uint `json:"id"`
}

type GetTaskResponse struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	BoardID     int      `json:"board_id"`
	ProjectID   int      `json:"project_id"`
	Order       int      `json:"order"`
	ParentID    int      `json:"parent_id"`
	Status      int      `json:"status"`
	Priority    int      `json:"priority"`
	DueDate     int      `json:"due_date"`
	AssigneeID  int      `json:"assignee"`
	Tags        []string `json:"tags"`
	CreatedBy   int      `json:"created_by"`
	CreatedAt   int      `json:"created_at"`
	UpdatedAt   int      `json:"updated_at"`
}

func (r *GetTaskResponse) FromModel(m *models.Tasks) *GetTaskResponse {
	res := &GetTaskResponse{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		BoardID:     m.BoardID,
		ProjectID:   m.ProjectID,
		Order:       m.Order,
		ParentID:    m.ParentID,
		Status:      m.Status,
		Tags:        m.Tags,
		CreatedBy:   m.CreatedBy,
		CreatedAt:   int(m.CreatedAt.Unix()),
		UpdatedAt:   int(m.UpdatedAt.Unix()),
	}
	if m.Priority != nil {
		res.Priority = *m.Priority
	}
	if m.DueDate != nil {
		res.DueDate = *m.DueDate
	}
	if m.AssigneeID != nil {
		res.AssigneeID = *m.AssigneeID
	}
	return res
}

type GetTasksResponse struct {
	Tasks []GetTaskResponse `json:"tasks"`
}

type CompleteTaskResponse struct {
	ID uint `json:"id"`
}

type DeleteTaskResponse struct {
	ID uint `json:"id"`
}
