package handlers

import (
	"fmt"
	"net/http"

	"github.com/imrinry/assessment-tax/errs"

	"github.com/labstack/echo/v4"
)

type AppResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func HandleError(c echo.Context, err error) error {
	switch e := err.(type) {
	case errs.AppError:
		fmt.Println("AppError: ", e)
		return c.JSON(e.Code, e)
	case error:
		fmt.Println("Error: ", e)
		return c.NoContent(http.StatusInternalServerError)
	default:
		fmt.Println("Default: ", e)
		return c.NoContent(http.StatusInternalServerError)
	}
}

func handleResponse(c echo.Context, d AppResponse) error {
	return c.JSON(d.Code, d)
}
