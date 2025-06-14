package repo

import (
	"errors"

	"github.com/Mhirii/TaskHub/backend/models"
	"gorm.io/gorm"
)

type ProjectsRepo interface {
	CreateProject(project models.Projects) (*models.Projects, error)
	GetProjectByID(projectID uint) (*models.Projects, error)
	GetProjectsByUserID(userID uint) ([]*models.Projects, error)
	UpdateProject(project *models.Projects) (*models.Projects, error)
	DeleteProject(project models.Projects) (*models.Projects, error)
	AddUser(projectID uint, userID uint) (*models.Projects, error)
	RemoveUser(projectID uint, userID uint) (*models.Projects, error)
}
type projectsRepo struct {
	db *gorm.DB
}

func NewProjectsRepo(db *gorm.DB) ProjectsRepo {
	return &projectsRepo{db: db}
}

func (r *projectsRepo) CreateProject(project models.Projects) (*models.Projects, error) {
	res := r.db.Create(&project)
	if res.Error != nil {
		return nil, res.Error
	}
	return &project, nil
}

func (r *projectsRepo) GetProjectByID(projectID uint) (*models.Projects, error) {
	var project models.Projects
	res := r.db.Where("id = ?", projectID).First(&project)
	if res.Error != nil {
		return nil, res.Error
	}

	return &project, nil
}

func (r *projectsRepo) GetProjectsByUserID(userID uint) ([]*models.Projects, error) {
	var projects []*models.Projects
	res := r.db.Where("owner_id = ?", userID).Find(&projects)
	if res.Error != nil {
		return nil, res.Error
	}

	return projects, nil
}

func (r *projectsRepo) UpdateProject(project *models.Projects) (*models.Projects, error) {
	res := r.db.Save(project)
	if res.Error != nil {
		return nil, res.Error
	}
	return project, nil
}

func (r *projectsRepo) DeleteProject(project models.Projects) (*models.Projects, error) {
	res := r.db.Delete(&project)
	if res.Error != nil {
		return nil, res.Error
	}
	return &project, nil
}

func (r *projectsRepo) AddUser(projectID uint, userID uint) (*models.Projects, error) {
	return nil, errors.New("not implemented")
}

func (r *projectsRepo) RemoveUser(projectID uint, userID uint) (*models.Projects, error) {
	return nil, errors.New("not implemented")
}
