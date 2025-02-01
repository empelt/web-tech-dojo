package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/empelt/web-tech-dojo/models"
	"github.com/labstack/echo/v4"
)

func NewQuestionHandler(authService AuthService, questionService QuestionService) (*QuestionHandler, error) {
	return &QuestionHandler{
		authService:     authService,
		questionService: questionService,
	}, nil
}

func (h *QuestionHandler) GetQuestion(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("id is required"))
	}
	qid, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("type error: id"))
	}

	uid, err := h.authService.AuthorizeAsUser(c.Request().Context(), getIdToken(c))
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	q, err := h.questionService.GetQuestion(c.Request().Context(), uid, qid)
	if err == models.EntityNotFoundError {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("internal server error"))
	}
	return c.JSON(http.StatusOK, q)
}

func (h *QuestionHandler) GetAllQuestions(c echo.Context) error {
	uid, err := h.authService.AuthorizeAsUser(c.Request().Context(), getIdToken(c))
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	qs, err := h.questionService.GetAllQuestions(c.Request().Context(), uid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, qs)
}
