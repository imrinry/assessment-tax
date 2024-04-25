package handlers

import (
	"net/http"

	"github.com/imrinry/assessment-tax/models"
	"github.com/labstack/echo/v4"
)

func (h *handler) TaxCalculations(c echo.Context) error {
	var req models.TaxRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	tax, err := h.s.TaxCalculations(c.Request().Context(), req.TotalIncome, req.WHT, req.Allowances)
	if err != nil {
		return HandleError(c, err)
	}
	resp := models.TaxResponse{Tax: tax}
	return c.JSON(http.StatusOK, resp)
}
