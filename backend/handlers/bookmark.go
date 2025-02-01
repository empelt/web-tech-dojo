package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func NewBookmarkHandler(as AuthService, bs BookmarkService) (*BookmarkHandler, error) {
	return &BookmarkHandler{
		authService:     as,
		bookmarkService: bs,
	}, nil
}

func (h *BookmarkHandler) GetBookmark(c echo.Context) error {
	uid, err := h.authService.AuthorizeAsUser(c.Request().Context(), getIdToken(c))
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err)
	}
	b, err := h.bookmarkService.GetBookmark(c.Request().Context(), uid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	qids := b.QuestionIds
	if qids == nil {
		qids = []int{}
	}
	return c.JSON(http.StatusOK, GetBookmarkResponse{
		QuestionIds: qids,
	})
}

func (h *BookmarkHandler) AddBookmark(c echo.Context) error {
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

	if err := h.bookmarkService.AddBookmark(c.Request().Context(), uid, qid); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return nil
}

func (h *BookmarkHandler) RemoveBookmark(c echo.Context) error {
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

	if err := h.bookmarkService.RemoveBookmark(c.Request().Context(), uid, qid); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return nil
}
