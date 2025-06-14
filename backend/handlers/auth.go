package handlers

import (
	"net/http"
	"strings"

	"github.com/Mhirii/TaskHub/backend/dto"
	"github.com/Mhirii/TaskHub/backend/services"
	"github.com/labstack/echo/v4"
)

type AuthHandlers struct {
	svc services.AuthService
}

func NewAuthHandlers(svc services.AuthService) *AuthHandlers {
	return &AuthHandlers{
		svc: svc,
	}
}

func (h *AuthHandlers) WriteGroup(g *echo.Group) {
	g.POST("/login", h.Login)
	g.POST("/register", h.Register)
	g.POST("/refresh", h.Refresh)
}

func (h *AuthHandlers) Login(c echo.Context) error {
	var body dto.LoginRequest
	err := c.Bind(&body)
	if err != nil {
		return err
	}
	resp, err := h.svc.Login(body.Username, body.Password)
	if err != nil {
		if strings.Contains(err.Error(), "found") {
			return c.JSON(http.StatusNotFound, err)
		}
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *AuthHandlers) Register(c echo.Context) error {
	var body dto.RegisterRequest
	err := c.Bind(&body)
	if err != nil {
		return err
	}
	resp, err := h.svc.Register(body.Username, body.Email, body.Password)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			if strings.Contains(err.Error(), "username") {
				return c.JSON(http.StatusBadRequest, "username already exists")
			}
			if strings.Contains(err.Error(), "email") {
				return c.JSON(http.StatusBadRequest, "email already exists")
			}
		}
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *AuthHandlers) Refresh(c echo.Context) error {
	var body dto.RefreshRequest
	err := c.Bind(&body)
	if err != nil {
		return err
	}
	resp, err := h.svc.Refresh(body.RefreshToken)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}
