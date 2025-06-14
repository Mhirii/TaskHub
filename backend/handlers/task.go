package handlers

import "github.com/labstack/echo/v4"

type TaskHandlers struct{}

func NewTaskHandlers() *TaskHandlers {
	return &TaskHandlers{}
}

func (h *TaskHandlers) WriteGroup(g *echo.Group) {}
