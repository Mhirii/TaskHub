package repo

import (
	"github.com/Mhirii/TaskHub/backend/models"
	"gorm.io/gorm"
)

type UsersRepo interface {
	CreateUser(username string, email string, password string) (*models.Users, error)
	GetUserByUsername(username string) (*models.Users, error)
	GetUserByEmail(email string) (*models.Users, error)
	UpdateUser(user *models.Users) (*models.Users, error)
	DeleteUser(userID uint) (*models.Users, error)
}
type usersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) UsersRepo {
	return &usersRepo{db: db}
}

func (r *usersRepo) CreateUser(username string, email string, password string) (*models.Users, error) {
	user := models.Users{Username: username, Email: email, Password: password}
	res := r.db.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (r *usersRepo) GetUserByUsername(username string) (*models.Users, error) {
	var user models.Users
	res := r.db.Where("username = ?", username).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (r *usersRepo) GetUserByEmail(email string) (*models.Users, error) {
	var user models.Users
	res := r.db.Where("email = ?", email).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (r *usersRepo) UpdateUser(user *models.Users) (*models.Users, error) {
	res := r.db.Save(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (r *usersRepo) DeleteUser(userID uint) (*models.Users, error) {
	user := models.Users{}
	user.ID = userID
	res := r.db.Delete(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
