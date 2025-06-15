package services

import (
	"errors"

	"github.com/Mhirii/TaskHub/backend/dto"
	"github.com/Mhirii/TaskHub/backend/pkg/roles"
	"github.com/Mhirii/TaskHub/backend/repo"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UsersService interface {
	CreateUser(ctx echo.Context, user dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	GetUserByID(ctx echo.Context, userID uint) (*dto.GetUserResponse, error)
	GetUserByUsername(ctx echo.Context, username string) (*dto.GetUserResponse, error)
	GetUsers(ctx echo.Context) ([]*dto.GetUserResponse, error)
	UpdateUser(ctx echo.Context, user dto.UpdateUserRequest) (*dto.UpdateUserResponse, error)
	DeleteUser(ctx echo.Context, userID uint) (*dto.DeleteUserResponse, error)
}

type usersService struct {
	repo repo.UsersRepo
}

func NewUsersService(repo repo.UsersRepo) UsersService {
	return &usersService{repo: repo}
}

func (s *usersService) CreateUser(ctx echo.Context, user dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	u := user.ToModel()
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	createdUser, err := s.repo.CreateUser(u.Username, u.Email, string(hashed), roles.StringToRoles(u.Roles))
	if err != nil {
		return nil, err
	}
	return &dto.CreateUserResponse{ID: createdUser.ID}, nil
}

func (s *usersService) GetUserByID(ctx echo.Context, userID uint) (*dto.GetUserResponse, error) {
	roles := ctx.Get("roles").([]string)
	requesterID := ctx.Get("userID").(uint)
	err := checkAccess(roles, userID, requesterID)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	res := &dto.GetUserResponse{}
	res = res.FromModel(user)
	return res, nil
}

func (s *usersService) GetUserByUsername(ctx echo.Context, username string) (*dto.GetUserResponse, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	res := &dto.GetUserResponse{}
	res = res.FromModel(user)
	return res, nil
}

func (s *usersService) GetUsers(ctx echo.Context) ([]*dto.GetUserResponse, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}
	var usersResponse []*dto.GetUserResponse
	for _, user := range users {
		res := &dto.GetUserResponse{}
		u := res.FromModel(user)
		usersResponse = append(usersResponse, u)
	}
	return usersResponse, nil
}

func (s *usersService) UpdateUser(ctx echo.Context, user dto.UpdateUserRequest) (*dto.UpdateUserResponse, error) {
	u := user.ToModel()
	updatedUser, err := s.repo.UpdateUser(u)
	if err != nil {
		return nil, err
	}
	return &dto.UpdateUserResponse{ID: updatedUser.ID}, nil
}

func (s *usersService) DeleteUser(ctx echo.Context, userID uint) (*dto.DeleteUserResponse, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	deletedUser, err := s.repo.DeleteUser(user.ID)
	if err != nil {
		return nil, err
	}
	return &dto.DeleteUserResponse{ID: deletedUser.ID}, nil
}

func checkAccess(roles []string, userID uint, requesterID uint) error {
	if requesterID == 0 {
		return errors.New("invalid requester")
	}
	if requesterID == userID {
		return nil
	}
	for _, role := range roles {
		if role == "admin" {
			return nil
		}
	}
	return errors.New("unauthorized")
}
