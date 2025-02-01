package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func NewAnswerHandler(as AuthService, cs AnswerService) (*AnswerHandler, error) {
	return &AnswerHandler{
		authService:   as,
		answerService: cs,
	}, nil
}

func getIdToken(c echo.Context) string {
	authorization := echo.Context.Request(c).Header.Get("Authorization")
	return strings.TrimPrefix(authorization, "Bearer ")
}

func (h *AnswerHandler) GetPreviousAnswer(c echo.Context) error {
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

	a, err := h.answerService.GetPreviousAnswer(c.Request().Context(), uid, qid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, BuildGetPreviousAnswerReponse(a))
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
	uid, err := h.authService.AuthorizeAsUser(c.Request().Context(), idToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	res, err := h.answerService.PostQuestionAnswer(c.Request().Context(), uid, qid, params.Answer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, PostQuestionAnswerResponse{Answer: res.Message})
}
