package services

import (
	"github.com/Mhirii/TaskHub/backend/dto"
	"github.com/Mhirii/TaskHub/backend/repo"
	"github.com/labstack/echo/v4"
)

type ProjectsService interface {
	CreateProject(c echo.Context, userID uint, project dto.CreateProjectRequest) (*dto.CreateProjectResponse, error)
	GetProjectByID(c echo.Context, projectID uint) (*dto.GetProjectResponse, error)
	GetProjects(c echo.Context, userID uint) ([]*dto.GetProjectResponse, error)
	UpdateProject(c echo.Context, projectID uint, project dto.UpdateProjectRequest) (*dto.UpdateProjectResponse, error)
	DeleteProject(c echo.Context, projectID uint) (*dto.DeleteProjectResponse, error)
	AddUser(c echo.Context, projectID uint, userID uint) (*dto.AddUserResponse, error)
	RemoveUser(c echo.Context, projectID uint, userID uint) (*dto.RemoveUserResponse, error)
}

type projectsService struct {
	repo repo.ProjectsRepo
}

func NewProjectsService(repo repo.ProjectsRepo) ProjectsService {
	return &projectsService{repo: repo}
}

func (s *projectsService) CreateProject(c echo.Context, userID uint, project dto.CreateProjectRequest) (*dto.CreateProjectResponse, error) {
	p := project.ToModel()
	p.OwnerID = int(userID)

	createdProject, err := s.repo.CreateProject(*p)
	if err != nil {
		return nil, err
	}
	return &dto.CreateProjectResponse{ID: createdProject.ID}, nil
}

func (s *projectsService) GetProjectByID(c echo.Context, projectID uint) (*dto.GetProjectResponse, error) {
	project, err := s.repo.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}
	res := &dto.GetProjectResponse{}
	res = res.FromModel(project)
	return res, nil
}

func (s *projectsService) GetProjects(c echo.Context, userID uint) ([]*dto.GetProjectResponse, error) {
	projects, err := s.repo.GetProjectsByUserID(userID)
	if err != nil {
		return nil, err
	}
	var projectsResponse []*dto.GetProjectResponse
	for _, project := range projects {
		res := &dto.GetProjectResponse{}
		res.FromModel(project)
		projectsResponse = append(projectsResponse, res)
	}
	return projectsResponse, nil
}

func (s *projectsService) UpdateProject(c echo.Context, projectID uint, project dto.UpdateProjectRequest) (*dto.UpdateProjectResponse, error) {
	p := project.ToModel()
	updatedProject, err := s.repo.UpdateProject(p)
	if err != nil {
		return nil, err
	}
	return &dto.UpdateProjectResponse{ID: updatedProject.ID}, nil
}

func (s *projectsService) DeleteProject(c echo.Context, projectID uint) (*dto.DeleteProjectResponse, error) {
	project, err := s.repo.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}
	deletedProject, err := s.repo.DeleteProject(*project)
	if err != nil {
		return nil, err
	}
	return &dto.DeleteProjectResponse{ID: deletedProject.ID}, nil
}

func (s *projectsService) AddUser(c echo.Context, projectID uint, userID uint) (*dto.AddUserResponse, error) {
	addedProject, err := s.repo.AddUser(projectID, userID)
	if err != nil {
		return nil, err
	}
	return &dto.AddUserResponse{ID: addedProject.ID}, nil
}

func (s *projectsService) RemoveUser(c echo.Context, projectID uint, userID uint) (*dto.RemoveUserResponse, error) {
	removedProject, err := s.repo.RemoveUser(projectID, userID)
	if err != nil {
		return nil, err
	}
	return &dto.RemoveUserResponse{ID: removedProject.ID}, nil
}
