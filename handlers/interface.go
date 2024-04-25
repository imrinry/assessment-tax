package handlers

import (
	"github.com/imrinry/assessment-tax/services"

	"github.com/labstack/echo/v4"
)

type Handlers interface {
	Index(c echo.Context) error
	ExamHandler(c echo.Context) error
}

type handler struct {
	s services.Services
}

func New(s services.Services) Handlers {
	return &handler{
		s: s,
	}
}
