package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/Mhirii/TaskHub/backend/dto"
	"github.com/Mhirii/TaskHub/backend/models"
	"github.com/Mhirii/TaskHub/backend/pkg/config"
	"github.com/Mhirii/TaskHub/backend/pkg/roles"
	"github.com/Mhirii/TaskHub/backend/repo"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"golang.org/x/crypto/bcrypt"
)

type TokensResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type ATokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type AuthService interface {
	Login(username string, password string) (*dto.AuthResponse, error)
	Register(username string, email string, password string) (*dto.AuthResponse, error)
	Refresh(refreshToken string) (*ATokenResponse, error)
}
type authService struct {
	repo repo.UsersRepo
}

func NewAuthService(repo repo.UsersRepo) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(username string, password string) (*dto.AuthResponse, error) {
	user, err := s.repo.GetUserByUserOrEmail(username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	access, refresh, exp, err := GenerateTokenPair(user)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		User:         dto.UserResponse{Username: user.Username, Email: user.Email},
		AccessToken:  access,
		RefreshToken: refresh,
		ExpiresIn:    exp,
	}, nil
}

func (s *authService) Register(username string, email string, password string) (*dto.AuthResponse, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.CreateUser(username, email, string(hashed), []string{})
	if err != nil {
		return nil, err
	}

	access, refresh, exp, err := GenerateTokenPair(user)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		User:         dto.UserResponse{Username: user.Username, Email: user.Email},
		AccessToken:  access,
		RefreshToken: refresh,
		ExpiresIn:    exp,
	}, nil
}

func (s *authService) Refresh(refreshToken string) (*ATokenResponse, error) {
	opts := []jwt.ParseOption{
		jwt.WithKey(jwa.HS512(), []byte(config.GetConfig().Auth.Secret)),
	}
	t, err := jwt.Parse([]byte(refreshToken), opts...)
	if err != nil {
		// TODO: match against jwt errors
		err := errors.New("Could not parse token: " + err.Error())
		return nil, err
	}
	sub, ok := t.Subject()
	if !ok {
		return nil, errors.New("could not get subject from token")
	}
	userID, err := strconv.Atoi(sub)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.GetUserByID(uint(userID))
	if err != nil {
		return nil, err
	}
	access, _, exp, err := GenerateTokenPair(user)
	if err != nil {
		return nil, err
	}

	return &ATokenResponse{
		AccessToken: access,
		ExpiresIn:   exp,
	}, nil
}

func JWTBuilder() *jwt.Builder {
	return jwt.NewBuilder().Issuer("taskhub").IssuedAt(time.Now())
}

func GenerateTokenPair(user *models.Users) (string, string, int, error) {
	id := strconv.Itoa(int(user.ID))
	access, err := GenerateToken(id, user.Username, roles.StringToRoles(user.Roles), config.GetConfig().Auth.AccessExp)
	if err != nil {
		return "", "", 0, err
	}
	refresh, err := GenerateToken(id, user.Username, roles.StringToRoles(user.Roles), config.GetConfig().Auth.RefreshExp)
	if err != nil {
		return "", "", 0, err
	}
	exp := time.Now().Add(time.Duration(config.GetConfig().Auth.AccessExp) * time.Second)
	return access, refresh, int(exp.Unix()), nil
}

func GenerateToken(id string, username string, roles []string, exp int) (string, error) {
	builder := JWTBuilder()
	cfg := config.GetConfig()
	expiry := time.Now().Add(time.Duration(exp) * time.Second)
	jwttoken, err := builder.
		Subject(string(id)).
		Expiration(expiry).
		Claim("username", username).
		Claim("roles", roles).
		Build()
	if err != nil {
		return "", err
	}
	opts := []jwt.SignOption{
		jwt.WithKey(jwa.HS512(), []byte(cfg.Auth.Secret)),
	}
	token, err := jwt.Sign(jwttoken, opts...)
	if err != nil {
		return "", err
	}
	return string(token), nil
}
