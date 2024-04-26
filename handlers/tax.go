package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gocarina/gocsv"
	"github.com/imrinry/assessment-tax/errs"
	"github.com/imrinry/assessment-tax/models"
	"github.com/labstack/echo/v4"
)

func (h *handler) TaxCalculations(c echo.Context) error {
	var req models.TaxRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	data, err := h.s.TaxCalculations(c.Request().Context(), req.TotalIncome, req.WHT, req.Allowances)
	if err != nil {
		return HandleError(c, err)
	}

	return c.JSON(http.StatusOK, data)
}
func (h *handler) CsvFileTaxCalculations(c echo.Context) error {

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	var allList []*models.ListFile
	if err := gocsv.Unmarshal(src, &allList); err != nil {
		return err
	}

	// validatedList := []models.ListForCalculateTax{}
	reqList := []models.TaxRequest{}
	for _, person := range allList {
		totalIncome, err := strconv.ParseFloat(person.TotalIncome, 64)
		if err != nil {
			HandleError(c, errs.NewValidationError(fmt.Sprintf("invalid totalIncome: %s", person.TotalIncome)))
		}
		wht, err := strconv.ParseFloat(person.Wht, 64)
		if err != nil {
			HandleError(c, errs.NewValidationError(fmt.Sprintf("invalid wht: %s", person.Wht)))
		}
		donation, err := strconv.ParseFloat(person.Donation, 64)
		if err != nil {
			HandleError(c, errs.NewValidationError(fmt.Sprintf("invalid donation: %s", person.Donation)))
		}
		// validatedList = append(validatedList, models.ListForCalculateTax{TotalIncome: totalIncome, Wht: wht, Donation: donation})
		reqList = append(reqList, models.TaxRequest{TotalIncome: totalIncome, WHT: wht, Allowances: []models.Allowance{models.Allowance{AllowanceType: models.DonationType, Amount: donation}}})
	}
	results := make([]map[string]interface{}, 0)
	for _, v := range reqList {
		r, err := h.s.TaxCalculations(c.Request().Context(), v.TotalIncome, v.WHT, v.Allowances)
		if err != nil {
			return HandleError(c, err)
		}
		result := map[string]interface{}{
			"totalIncome": v.TotalIncome,
			"tax":         r.Tax,
		}
		results = append(results, result)
	}
	res := make(map[string]interface{})
	res["taxes"] = results
	return c.JSON(http.StatusOK, res)
}
