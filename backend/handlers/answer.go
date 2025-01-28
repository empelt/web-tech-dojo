package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
)

type PostQuestionAnswerRequest struct {
	Answer string `json:"message" validate:"required"`
}

type PostQuestionAnswerResponse struct {
	Answer string `json:"message"`
}

func NewAnswerHandler(authClient auth.Client, cs AnswerService) (*AnswerHandler, error) {
	return &AnswerHandler{
		authClient:    &authClient,
		answerService: cs,
	}, nil
}

func (h *AnswerHandler) PostQuestionAnswer(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("id is required"))
	}
	qid, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("type error: id"))
	}

	params := &PostQuestionAnswerRequest{}
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	idToken := echo.Context.Request(c).Header.Get("Authorization")
	token, err := h.authClient.VerifyIDToken(c.Request().Context(), idToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	res, err := h.answerService.PostQuestionAnswer(c.Request().Context(), token.UID, qid, params.Answer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, PostQuestionAnswerResponse{Answer: res.Message})
}
