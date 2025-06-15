package dto

import (
	"strings"

	"github.com/Mhirii/TaskHub/backend/models"
	"github.com/jinzhu/copier"
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
		Tags:        strings.Join(r.Tags, "--"),
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
	ID          uint     `json:"id"`
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
	var task models.Tasks
	err := copier.Copy(&task, r)
	if err != nil {
		return nil
	}
	return &task
}

type UpdateTaskResponse struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	BoardID     int      `json:"board_id,omitempty"`
	ProjectID   int      `json:"project_id,omitempty"`
	Order       int      `json:"order,omitempty"`
	ParentID    int      `json:"parent_id,omitempty"`
	Status      int      `json:"status,omitempty"`
	Priority    int      `json:"priority,omitempty"`
	DueDate     int      `json:"due_date,omitempty"`
	AssigneeID  int      `json:"assignee,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	CreatedBy   int      `json:"created_by,omitempty"`
	CreatedAt   int      `json:"created_at,omitempty"`
	UpdatedAt   int      `json:"updated_at,omitempty"`
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
	r.ID = m.ID
	r.Name = m.Name
	r.Description = m.Description
	r.BoardID = m.BoardID
	r.ProjectID = m.ProjectID
	r.Order = m.Order
	r.ParentID = m.ParentID
	r.Status = m.Status
	r.Tags = strings.Split(m.Tags, "--")
	r.CreatedBy = m.CreatedBy
	r.CreatedAt = int(m.CreatedAt.Unix())
	r.UpdatedAt = int(m.UpdatedAt.Unix())
	if m.Priority != nil {
		r.Priority = *m.Priority
	}
	if m.DueDate != nil {
		r.DueDate = *m.DueDate
	}
	if m.AssigneeID != nil {
		r.AssigneeID = *m.AssigneeID
	}
	return r
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

type TaskFilters struct {
	ProjectID *int `json:"project_id"`
	Status    *int `json:"status"`
}
