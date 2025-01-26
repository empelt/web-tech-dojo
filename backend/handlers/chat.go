package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostQuestionAnswerRequest struct {
	Answer string `json:"message" validate:"required"`
}

type PostQuestionAnswerResponse struct {
	Answer string `json:"message"`
}

func New(cs ChatService) (*ChatHandler, error) {
	return &ChatHandler{
		chatService: cs,
	}, nil
}

func (h *ChatHandler) PostQuestionAnswer(c echo.Context) error {
	params := &PostQuestionAnswerRequest{}
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := h.chatService.PostQuestionAnswer(c.Request().Context(), params.Answer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, PostQuestionAnswerResponse{Answer: res})
}
