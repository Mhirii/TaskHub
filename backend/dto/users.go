package dto

import (
	"github.com/Mhirii/TaskHub/backend/models"
	"github.com/Mhirii/TaskHub/backend/pkg/roles"
)

type CreateUserRequest struct {
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

func (r *CreateUserRequest) ToModel() *models.Users {
	user := &models.Users{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		Roles:    roles.RolesToString(r.Roles),
	}
	return user
}

type CreateUserResponse struct {
	ID uint `json:"id"`
}

type UpdateUserRequest struct {
	Username *string   `json:"username"`
	Email    *string   `json:"email"`
	Password *string   `json:"password"`
	Roles    *[]string `json:"roles"`
}

func (r *UpdateUserRequest) ToModel() *models.Users {
	user := &models.Users{}
	if r.Roles != nil {
		user.Roles = roles.RolesToString(*r.Roles)
	}
	if r.Username != nil {
		user.Username = *r.Username
	}
	if r.Email != nil {
		user.Email = *r.Email
	}
	if r.Password != nil {
		user.Password = *r.Password
	}
	return user
}

type UpdateUserResponse struct {
	ID uint `json:"id"`
}

type UserProject struct {
	ProjectID     uint   `json:"project_id"`
	Name          string `json:"name"`
	Role          string `json:"role"`
	TasksAssigned int    `json:"tasks_assigned"`
}

type GetUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`

	Roles    []string      `json:"roles"`
	Projects []UserProject `json:"projects"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (r *GetUserResponse) FromModel(m *models.Users) *GetUserResponse {
	res := &GetUserResponse{
		ID:        m.ID,
		Username:  m.Username,
		Email:     m.Email,
		CreatedAt: m.CreatedAt.String(),
		UpdatedAt: m.UpdatedAt.String(),
		Projects:  []UserProject{},
		Roles:     roles.StringToRoles(m.Roles),
	}
	r = res
	return res
}

type GetUsersResponse struct {
	Users []GetUserResponse `json:"users"`
}

type DeleteUserResponse struct {
	ID uint `json:"id"`
}
