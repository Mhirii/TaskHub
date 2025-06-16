package services

import (
	"github.com/Mhirii/TaskHub/backend/models"
	"github.com/Mhirii/TaskHub/backend/repo"
	"gorm.io/gorm"
)

type UserProjectsService struct {
	userProjectsRepo repo.UserProjectsRepo
}

func NewUserProjectsService(userProjectsRepo repo.UserProjectsRepo) *UserProjectsService {
	return &UserProjectsService{userProjectsRepo: userProjectsRepo}
}

// add user to project
// remote user from project
// change role of user in project
// get all user projects by userid, and by projectid, and by userid or projectid
func (s *UserProjectsService) CreateUserProject(userID uint, projectID uint, role string) (*models.UserProjects, error) {
	userProject, err := s.userProjectsRepo.CreateUserProject(userID, projectID, role)
	if err != nil {
		return nil, err
	}
	return userProject, nil
}

func (s *UserProjectsService) GetUserProjectsByUserID(userID uint) ([]*models.UserProjects, error) {
	userProjects, err := s.userProjectsRepo.GetUserProjectsByUserID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*models.UserProjects{}, nil
		}
		return nil, err
	}
	return userProjects, nil
}

func (s *UserProjectsService) GetUserProjectsByProjectID(projectID uint) ([]*models.UserProjects, error) {
	userProjects, err := s.userProjectsRepo.GetUserProjectsByProjectID(projectID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*models.UserProjects{}, nil
		}
		return nil, err
	}
	return userProjects, nil
}

func (s *UserProjectsService) GetUserProjectsByUserOrProjectID(userOrProjectID string) ([]*models.UserProjects, error) {
	userProjects, err := s.userProjectsRepo.GetUserProjectsByUserOrProjectID(userOrProjectID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*models.UserProjects{}, nil
		}
		return nil, err
	}
	return userProjects, nil
}

func (s *UserProjectsService) DeleteUserProject(userID uint, projectID uint) error {
	if err := s.userProjectsRepo.DeleteUserProject(userID, projectID); err != nil {
		return err
	}
	return nil
}

func (s *UserProjectsService) UpdateUserProjectRole(userID uint, projectID uint, newRole string) error {
	if err := s.userProjectsRepo.UpdateUserProjectRole(userID, projectID, newRole); err != nil {
		return err
	}
	return nil
}
