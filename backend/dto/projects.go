package dto

import (
	"github.com/Mhirii/TaskHub/backend/models"
)

type CreateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateProjectRequest) ToModel() *models.Projects {
	project := &models.Projects{
		Name:        r.Name,
		Description: r.Description,
	}
	return project
}

type CreateProjectResponse struct {
	ID uint `json:"id"`
}

type UpdateProjectRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	OwnerID     *int    `json:"owner_id"`
}

func (r *UpdateProjectRequest) ToModel() *models.Projects {
	project := &models.Projects{}
	if r.OwnerID != nil {
		project.OwnerID = *r.OwnerID
	}
	if r.Name != nil {
		project.Name = *r.Name
	}
	if r.Description != nil {
		project.Description = *r.Description
	}
	return project
}

type UpdateProjectResponse struct {
	ID uint `json:"id"`
}

type GetProjectResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     int    `json:"owner_id"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
}

func (r *GetProjectResponse) FromModel(m *models.Projects) *GetProjectResponse {
	r.ID = m.ID
	r.Name = m.Name
	r.Description = m.Description
	r.OwnerID = m.OwnerID
	r.CreatedAt = int(m.CreatedAt.Unix())
	r.UpdatedAt = int(m.UpdatedAt.Unix())
	return r
}

type GetProjectsResponse struct {
	Projects []GetProjectResponse `json:"projects"`
}

type AddUserResponse struct {
	ID uint `json:"id"`
}

type RemoveUserResponse struct {
	ID uint `json:"id"`
}

type DeleteProjectResponse struct {
	ID uint `json:"id"`
}
