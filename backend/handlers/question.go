package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

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
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.New("internal server error"))
	}
	if q.Question == nil {
		return echo.NewHTTPError(http.StatusNotFound, errors.New("question is not found"))
	}
	return c.JSON(http.StatusOK, &GetQuestionResponse{
		Id:        q.Question.Id,
		Title:     q.Question.Title,
		Content:   q.Question.Content,
		Tags:      q.Question.Tags,
		CreatedAt: q.Question.CreatedAt,
	})
}
