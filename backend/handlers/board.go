package handlers

import (
	"github.com/labstack/echo/v4"
)

type BoardHandlers struct {
	// svc services.BoardService
}

func NewBoardHandlers() *BoardHandlers {
	return &BoardHandlers{}
}

func (h *BoardHandlers) WriteGroup(g *echo.Group) {
}
