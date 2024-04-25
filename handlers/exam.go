package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) Index(c echo.Context) error {
	return handleResponse(c, AppResponse{Message: "success ", Code: http.StatusOK, Data: nil})
}

func (h *handler) ExamHandler(c echo.Context) error {
	data, err := h.s.ExamServices(c.Request().Context())
	if err != nil {
		return HandleError(c, err)
	}
	return handleResponse(c, AppResponse{Message: "success ", Code: http.StatusOK, Data: data})
}
