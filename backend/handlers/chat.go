package handlers

import (
	"net/http"

	"github.com/empelt/web-tech-dojo/models"
	"github.com/labstack/echo/v4"
)

func New(cs ChatService) (*ChatHandler, error) {
	return &ChatHandler{
		chatService: cs,
	}, nil
}

func (h *ChatHandler) PostChatMessage(c echo.Context) error {
	params := &models.PostChatMessageRequest{}
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := h.chatService.PostChatMessage(c.Request().Context(), params.Message)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, models.PostChatMessageResponse{Message: res})
}
