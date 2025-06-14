package services

import (
	"github.com/Mhirii/TaskHub/backend/dto"
	"github.com/Mhirii/TaskHub/backend/repo"
	"github.com/labstack/echo/v4"
)

type TaskService interface {
	CreateTask(c echo.Context, userID uint, projectID uint, task dto.CreateTaskRequest) (*dto.CreateTaskResponse, error)
	GetTaskByID(c echo.Context, taskID uint) (*dto.GetTaskResponse, error)
	GetTasks(c echo.Context, projectID uint) ([]*dto.GetTaskResponse, error)
	UpdateTask(c echo.Context, userID uint, task dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, error)
	CompleteTask(c echo.Context, userID uint, taskID uint) (*dto.CompleteTaskResponse, error)
	DeleteTask(c echo.Context, userID uint, taskID uint) (*dto.DeleteTaskResponse, error)
}

type taskService struct {
	repo repo.TasksRepo
}

func NewTasksService(repo repo.TasksRepo) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(c echo.Context, userID uint, projectID uint, task dto.CreateTaskRequest) (*dto.CreateTaskResponse, error) {
	t := task.ToModel()
	t.CreatedBy = int(userID)
	t.ProjectID = int(projectID)

	createdTask, err := s.repo.CreateTask(*t)
	if err != nil {
		return nil, err
	}
	return &dto.CreateTaskResponse{ID: createdTask.ID}, nil
}

func (s *taskService) GetTaskByID(c echo.Context, taskID uint) (*dto.GetTaskResponse, error) {
	task, err := s.repo.GetTaskByID(taskID)
	if err != nil {
		return nil, err
	}
	res := &dto.GetTaskResponse{}
	res = res.FromModel(task)
	return res, nil
}

func (s *taskService) GetTasks(c echo.Context, projectID uint) ([]*dto.GetTaskResponse, error) {
	tasks, err := s.repo.GetTasksByProjectID(projectID)
	if err != nil {
		return nil, err
	}
	var tasksResponse []*dto.GetTaskResponse
	for _, task := range tasks {
		res := &dto.GetTaskResponse{}
		res.FromModel(task)
		tasksResponse = append(tasksResponse, res)
	}
	return tasksResponse, nil
}

func (s *taskService) UpdateTask(c echo.Context, userID uint, task dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, error) {
	t := task.ToModel()
	updatedTask, err := s.repo.UpdateTask(t)
	if err != nil {
		return nil, err
	}
	return &dto.UpdateTaskResponse{ID: updatedTask.ID}, nil
}

func (s *taskService) CompleteTask(c echo.Context, userID uint, taskID uint) (*dto.CompleteTaskResponse, error) {
	task, err := s.repo.GetTaskByID(taskID)
	if err != nil {
		return nil, err
	}
	task.Status = 1
	updatedTask, err := s.repo.UpdateTask(task)
	if err != nil {
		return nil, err
	}
	return &dto.CompleteTaskResponse{ID: updatedTask.ID}, nil
}

func (s *taskService) DeleteTask(c echo.Context, userID uint, taskID uint) (*dto.DeleteTaskResponse, error) {
	task, err := s.repo.GetTaskByID(taskID)
	if err != nil {
		return nil, err
	}
	deletedTask, err := s.repo.DeleteTask(*task)
	if err != nil {
		return nil, err
	}
	return &dto.DeleteTaskResponse{ID: deletedTask.ID}, nil
}
