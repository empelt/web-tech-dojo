package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(api *echo.Group) {

	chat := api.Group("/chat")
	chat.POST("", h.PostChatMessage)
}
