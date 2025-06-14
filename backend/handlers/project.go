package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Mhirii/TaskHub/backend/dto"
	"github.com/Mhirii/TaskHub/backend/middleware"
	"github.com/Mhirii/TaskHub/backend/services"
	"github.com/labstack/echo/v4"
)

type ProjectHandlers struct {
	svc services.ProjectsService
}

func NewProjectHandlers(svc services.ProjectsService) *ProjectHandlers {
	return &ProjectHandlers{
		svc: svc,
	}
}

func (h *ProjectHandlers) WriteGroup(g *echo.Group) {
	g.Use(middleware.AuthMiddleware)
	g.POST("/", h.CreateProject)
	g.GET("/:id", h.GetProject)
	g.GET("/", h.GetProjects)
	g.PATCH("/:id", h.UpdateProject)
	g.DELETE("/:id", h.DeleteProject)
	g.POST("/:id/user", h.AddUser)
	g.DELETE("/:id/user", h.RemoveUser)
}

func (h *ProjectHandlers) CreateProject(c echo.Context) error {
	userIDstr := c.Get("userID").(string)
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		return err
	}
	var body dto.CreateProjectRequest
	if err := c.Bind(&body); err != nil {
		return err
	}
	resp, err := h.svc.CreateProject(c, uint(userID), body)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			if strings.Contains(err.Error(), "name") {
				return c.JSON(http.StatusBadRequest, "name already exists")
			}
			return c.JSON(http.StatusBadRequest, "a field already exists")
		}
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *ProjectHandlers) GetProject(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	resp, err := h.svc.GetProjectByID(c, uint(projectID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, err)
		}
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *ProjectHandlers) GetProjects(c echo.Context) error {
	fmt.Println("get projects")
	userIDstr := c.Get("userID").(string)
	fmt.Println("userID")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		return err
	}
	fmt.Println("userID: ", userID)
	resp, err := h.svc.GetProjects(c, uint(userID))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *ProjectHandlers) UpdateProject(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	var body dto.UpdateProjectRequest
	if err := c.Bind(&body); err != nil {
		return err
	}
	resp, err := h.svc.UpdateProject(c, uint(projectID), body)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *ProjectHandlers) DeleteProject(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	resp, err := h.svc.DeleteProject(c, uint(projectID))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *ProjectHandlers) AddUser(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	userIDstr := c.Get("userID").(string)
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		return err
	}
	resp, err := h.svc.AddUser(c, uint(projectID), uint(userID))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *ProjectHandlers) RemoveUser(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	userIDstr := c.Get("userID").(string)
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		return err
	}
	resp, err := h.svc.RemoveUser(c, uint(projectID), uint(userID))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}
