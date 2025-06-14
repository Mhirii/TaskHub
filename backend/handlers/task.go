package handlers

import (
	"net/http"
	"strconv"

	"github.com/Mhirii/TaskHub/backend/dto"
	"github.com/Mhirii/TaskHub/backend/middleware"
	"github.com/Mhirii/TaskHub/backend/services"
	"github.com/labstack/echo/v4"
)

type TaskHandlers struct {
	svc services.TaskService
}

func NewTaskHandlers(svc services.TaskService) *TaskHandlers {
	return &TaskHandlers{
		svc: svc,
	}
}

func (h *TaskHandlers) WriteGroup(g *echo.Group) {
	g.Use(middleware.AuthMiddleware)
	g.POST("/", h.CreateTask)
	g.GET("/:id", h.GetTask)
	g.GET("/", h.GetTasks)
	g.PATCH("/:id", h.UpdateTask)
	g.DELETE("/delete", h.DeleteTask)
	g.PATCH("/:id/complete", h.CompleteTask)
}

func (h *TaskHandlers) CreateTask(c echo.Context) error {
	userIDstr := c.Get("userID").(string)
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		return err
	}
	var body dto.CreateTaskRequest
	if err := c.Bind(&body); err != nil {
		return err
	}
	resp, err := h.svc.CreateTask(uint(userID), uint(body.ProjectID), body)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *TaskHandlers) GetTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	resp, err := h.svc.GetTaskByID(uint(taskID))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *TaskHandlers) GetTasks(c echo.Context) error {
	return nil
}

func (h *TaskHandlers) UpdateTask(c echo.Context) error {
	return nil
}

func (h *TaskHandlers) DeleteTask(c echo.Context) error {
	return nil
}

func (h *TaskHandlers) CompleteTask(c echo.Context) error {
	return nil
}
