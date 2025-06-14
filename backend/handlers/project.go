package handlers

import "github.com/labstack/echo/v4"

type ProjectHandlers struct{}

func NewProjectHandlers() *ProjectHandlers {
	return &ProjectHandlers{}
}

func (h *ProjectHandlers) WriteGroup(g *echo.Group) {}
