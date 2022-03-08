package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) userHello(c echo.Context) error {
	user := h.repo.getUser()
	v := struct {
		Msg string `json:"msg"`
	}{
		Msg: fmt.Sprintf("Hello, %s", user.Name),
	}
	return c.JSON(http.StatusOK, v)
}
