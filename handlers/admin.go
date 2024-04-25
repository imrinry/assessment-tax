package handlers

import (
	"net/http"

	"github.com/imrinry/assessment-tax/models"
	"github.com/labstack/echo/v4"
)

func (h *handler) DeductionPersonalSetting(c echo.Context) error {
	var req models.AdminSettings
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	_, err := h.s.DeductionPersonalSetting(c.Request().Context(), req.Amount)
	if err != nil {
		return HandleError(c, err)
	}
	res := make(map[string]interface{})
	res["personalDeduction"] = req.Amount
	return c.JSON(http.StatusOK, res)
}
func (h *handler) DeductionkReceiptSetting(c echo.Context) error {
	var req models.AdminSettings
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	_, err := h.s.DeductionkReceiptSetting(c.Request().Context(), req.Amount)
	if err != nil {
		return HandleError(c, err)
	}
	res := make(map[string]interface{})
	res["kReceipt"] = req.Amount
	return c.JSON(http.StatusOK, res)
}
