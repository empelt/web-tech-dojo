package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/empelt/web-tech-dojo/models"
	"github.com/labstack/echo/v4"
)

type GetQuestionResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewQuestionHandler(questionService QuestionService) (*QuestionHandler, error) {
	return &QuestionHandler{
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

	q, err := h.questionService.GetQuestion(c.Request().Context(), qid)
	if err == models.EntityNotFoundError {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("internal server error"))
	}
	return c.JSON(http.StatusOK, q)
}

func (h *QuestionHandler) GetAllQuestions(c echo.Context) error {
	qs, err := h.questionService.GetAllQuestions(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, qs)
}
