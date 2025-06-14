package services

import (
	"github.com/Mhirii/TaskHub/backend/dto"
	"github.com/Mhirii/TaskHub/backend/repo"
)

type TaskService interface {
	CreateTask(userID uint, projectID uint, task dto.CreateTaskRequest) (*dto.CreateTaskResponse, error)
	GetTaskByID(taskID uint) (*dto.GetTaskResponse, error)
	GetTasks(userID uint, projectID uint) ([]*dto.GetTaskResponse, error)
	UpdateTask(userID uint, task dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, error)
	CompleteTask(userID uint, taskID uint) (*dto.CompleteTaskResponse, error)
	DeleteTask(userID uint, taskID uint) (*dto.DeleteTaskResponse, error)
}

type taskService struct {
	repo repo.TasksRepo
}

func NewTasksService(repo repo.TasksRepo) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(userID uint, projectID uint, task dto.CreateTaskRequest) (*dto.CreateTaskResponse, error) {
	t := task.ToModel()
	t.CreatedBy = int(userID)
	t.ProjectID = int(projectID)

	createdTask, err := s.repo.CreateTask(*t)
	if err != nil {
		return nil, err
	}
	return &dto.CreateTaskResponse{ID: createdTask.ID}, nil
}

func (s *taskService) GetTaskByID(taskID uint) (*dto.GetTaskResponse, error) {
	task, err := s.repo.GetTaskByID(taskID)
	if err != nil {
		return nil, err
	}
	res := &dto.GetTaskResponse{}
	res = res.FromModel(task)
	return res, nil
}

func (s *taskService) GetTasks(userID uint, projectID uint) ([]*dto.GetTaskResponse, error) {
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

func (s *taskService) UpdateTask(userID uint, task dto.UpdateTaskRequest) (*dto.UpdateTaskResponse, error) {
	t := task.ToModel()
	updatedTask, err := s.repo.UpdateTask(t)
	if err != nil {
		return nil, err
	}
	return &dto.UpdateTaskResponse{ID: updatedTask.ID}, nil
}

func (s *taskService) CompleteTask(userID uint, taskID uint) (*dto.CompleteTaskResponse, error) {
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

func (s *taskService) DeleteTask(userID uint, taskID uint) (*dto.DeleteTaskResponse, error) {
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
