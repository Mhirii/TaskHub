package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Mhirii/TaskHub/backend/dto"
	"github.com/Mhirii/TaskHub/backend/middleware"
	"github.com/Mhirii/TaskHub/backend/services"
	"github.com/labstack/echo/v4"
)

type UsersHandlers struct {
	svc services.UsersService
}

func NewUsersHandlers(svc services.UsersService) *UsersHandlers {
	return &UsersHandlers{
		svc: svc,
	}
}

func (h *UsersHandlers) WriteGroup(g *echo.Group) {
	role := "admin"
	permission := "manage_users"
	g.Use(middleware.AuthMiddleware)
	g.Use(middleware.AccessMiddleware([]middleware.AccessMiddlewareRule{
		{
			AllowSelf: true,
		},
		{
			Role: &role,
		},
		{
			Permission: &permission,
		},
	}))
	g.POST("/", h.CreateUser)
	g.GET("/", h.GetUsers)
	g.GET("/:id", h.GetUser)
	g.PATCH("/:id", h.UpdateUser)
	g.DELETE("/:id", h.DeleteUser)
}

func (h *UsersHandlers) CreateUser(c echo.Context) error {
	roles := c.Get("roles").([]string)
	if len(roles) == 0 {
		return c.JSON(http.StatusBadRequest, "no roles")
	}
	hasAdmin := false
	for _, role := range roles {
		if role == "admin" {
			hasAdmin = true
			break
		}
	}
	if !hasAdmin {
		return c.JSON(http.StatusUnauthorized, "user must have the admin role")
	}
	var body dto.CreateUserRequest
	err := c.Bind(&body)
	if err != nil {
		return err
	}
	resp, err := h.svc.CreateUser(c, body)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, resp)
}

func (h *UsersHandlers) GetUser(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return err
	}
	res, err := h.svc.GetUserByID(c, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "unauthorized") {
			return c.JSON(http.StatusUnauthorized, err)
		}
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *UsersHandlers) GetUsers(c echo.Context) error {
	res, err := h.svc.GetUsers(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *UsersHandlers) UpdateUser(c echo.Context) error {
	res, err := h.svc.UpdateUser(c, dto.UpdateUserRequest{})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *UsersHandlers) DeleteUser(c echo.Context) error {
	res, err := h.svc.DeleteUser(c, 0)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
