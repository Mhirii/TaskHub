package repo

import (
	"fmt"
	"time"

	"github.com/Mhirii/TaskHub/backend/models"
	"gorm.io/gorm"
)

type TasksRepo interface {
	CreateTask(task models.Tasks) (*models.Tasks, error)
	GetTaskByID(taskID uint) (*models.Tasks, error)
	GetTasksByProjectID(projectID uint) ([]*models.Tasks, error)
	UpdateTask(taskID uint, task *models.Tasks) (*models.Tasks, error)
	DeleteTask(task models.Tasks) (*models.Tasks, error)
}
type tasksRepo struct {
	db *gorm.DB
}

func NewTasksRepo(db *gorm.DB) TasksRepo {
	return &tasksRepo{db: db}
}

func (r *tasksRepo) CreateTask(task models.Tasks) (*models.Tasks, error) {
	res := r.db.Create(&task)
	if res.Error != nil {
		return nil, res.Error
	}
	return &task, nil
}

func (r *tasksRepo) GetTaskByID(taskID uint) (*models.Tasks, error) {
	var task models.Tasks
	res := r.db.Where("id = ?", taskID).First(&task)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return new(models.Tasks), nil
		}
		return nil, res.Error
	}

	return &task, nil
}

func (r *tasksRepo) GetTasksByProjectID(projectID uint) ([]*models.Tasks, error) {
	var tasks []*models.Tasks
	res := r.db.Where("project_id = ?", projectID).Find(&tasks)
	if res.Error != nil {
		fmt.Println(res.Error)
		return nil, res.Error
	}

	return tasks, nil
}

func (r *tasksRepo) UpdateTask(id uint, task *models.Tasks) (*models.Tasks, error) {
	task.UpdatedAt = time.Now()
	fmt.Println(task.Description)
	res := r.db.Model(&task).Where("id = ?", id).Updates(task)
	if res.Error != nil {
		return nil, res.Error
	}
	return task, nil
}

func (r *tasksRepo) DeleteTask(task models.Tasks) (*models.Tasks, error) {
	res := r.db.Delete(&task)
	if res.Error != nil {
		return nil, res.Error
	}
	return &task, nil
}
