package repo

import (
	"errors"

	"github.com/Mhirii/TaskHub/backend/models"
	"gorm.io/gorm"
)

type UsersRepo interface {
	CreateUser(username string, email string, password string) (*models.Users, error)
	GetUserByUsername(username string) (*models.Users, error)
	GetUserByEmail(email string) (*models.Users, error)
	UpdateUser(user *models.Users) error
	DeleteUser(user *models.Users) error
}
type usersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) UsersRepo {
	return &usersRepo{db: db}
}

func (r *usersRepo) CreateUser(username string, email string, password string) (*models.Users, error) {
	return nil, errors.New("not implemented")
}

func (r *usersRepo) GetUserByUsername(username string) (*models.Users, error) {
	return nil, errors.New("not implemented")
}

func (r *usersRepo) GetUserByEmail(email string) (*models.Users, error) {
	return nil, errors.New("not implemented")
}

func (r *usersRepo) UpdateUser(user *models.Users) error {
	return errors.New("not implemented")
}

func (r *usersRepo) DeleteUser(user *models.Users) error {
	return errors.New("not implemented")
}
