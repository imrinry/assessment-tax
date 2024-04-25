package services

import (
	"context"

	"github.com/imrinry/assessment-tax/models"
)

func (s *service) TaxCalculations(ctx context.Context, income float64, wht float64, allowances []models.Allowance) (float64, error) {

	taxableIncome := income
	taxableIncome -= 60000
	type TaxLevel struct {
		Amount float64 `json:"amount"`
	}

	var totalTax float64
	var taxLevels []TaxLevel

	taxLevels = append(taxLevels, TaxLevel{})
	if taxableIncome > 150000 {
		nextLevel := min(taxableIncome, 500000) - 150000
		taxLevels = append(taxLevels, TaxLevel{Amount: nextLevel * 0.10})
	}
	if taxableIncome > 500000 {
		nextLevel := min(taxableIncome, 1000000) - 500000
		taxLevels = append(taxLevels, TaxLevel{Amount: nextLevel * 0.15})
	}
	if taxableIncome > 1000000 {
		nextLevel := min(taxableIncome, 2000000) - 1000000
		taxLevels = append(taxLevels, TaxLevel{Amount: nextLevel * 0.20})
	}
	if taxableIncome > 2000000 {
		nextLevel := taxableIncome - 2000000
		taxLevels = append(taxLevels, TaxLevel{Amount: nextLevel * 0.35})
	}

	for _, level := range taxLevels {
		totalTax += level.Amount
	}
	// totalTax -= wht
	return totalTax, nil
}
