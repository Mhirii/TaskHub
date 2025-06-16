package repo

import (
	"github.com/Mhirii/TaskHub/backend/models"
	"gorm.io/gorm"
)

type UserProjectsRepo struct {
	db *gorm.DB
}

func NewUserProjectsRepo(db *gorm.DB) *UserProjectsRepo {
	return &UserProjectsRepo{db: db}
}

func (r *UserProjectsRepo) CreateUserProject(userID uint, projectID uint, role string) (*models.UserProjects, error) {
	userProject := models.UserProjects{UserID: userID, ProjectID: projectID, Role: role}
	res := r.db.Create(&userProject)
	if res.Error != nil {
		return nil, res.Error
	}
	return &userProject, nil
}

func (r *UserProjectsRepo) GetUserProjectsByUserID(userID uint) ([]*models.UserProjects, error) {
	var userProjects []*models.UserProjects
	res := r.db.Where("user_id = ?", userID).Find(&userProjects)
	if res.Error != nil {
		return nil, res.Error
	}
	return userProjects, nil
}

func (r *UserProjectsRepo) GetUserProjectsByProjectID(projectID uint) ([]*models.UserProjects, error) {
	var userProjects []*models.UserProjects
	res := r.db.Where("project_id = ?", projectID).Find(&userProjects)
	if res.Error != nil {
		return nil, res.Error
	}
	return userProjects, nil
}

func (r *UserProjectsRepo) GetUserProjectsByUserOrProjectID(userOrProjectID string) ([]*models.UserProjects, error) {
	var userProjects []*models.UserProjects
	res := r.db.Where("user_id = ? OR project_id = ?", userOrProjectID, userOrProjectID).Find(&userProjects)
	if res.Error != nil {
		return nil, res.Error
	}
	return userProjects, nil
}

func (r *UserProjectsRepo) DeleteUserProject(userID uint, projectID uint) error {
	res := r.db.Where("user_id = ? AND project_id = ?", userID, projectID).Delete(&models.UserProjects{})
	return res.Error
}

func (r *UserProjectsRepo) UpdateUserProjectRole(userID uint, projectID uint, newRole string) error {
	res := r.db.Model(&models.UserProjects{}).Where("user_id = ? AND project_id = ?", userID, projectID).Update("role", newRole)
	return res.Error
}
