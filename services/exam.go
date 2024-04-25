package services

import (
	"context"
	"fmt"

	"github.com/imrinry/assessment-tax/errs"
	"github.com/imrinry/assessment-tax/models"
)

func (s *service) TaxCalculations(ctx context.Context, income float64, wht float64, allowances []models.Allowance) (models.TaxResponse, error) {

	taxableIncome := income

	if err := IsValidAllowanceValue(models.WhtType, wht, taxableIncome); err != nil {
		return models.TaxResponse{}, err
	}

	err := DeductTaxAllowances(allowances, &taxableIncome)
	if err != nil {
		return models.TaxResponse{}, err

	}
	taxableIncome -= models.PersonalDeductionValue

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
	totalTax -= wht
	fmt.Println("totalTax: ", totalTax)
	if totalTax < 0 {
		return models.TaxResponse{Tax: 0, TaxRefund: totalTax * -1}, nil
	}

	return models.TaxResponse{Tax: totalTax}, nil
}

func DeductTaxAllowances(allowances []models.Allowance, taxableIncome *float64) error {
	for _, allowance := range allowances {
		if err := IsValidAllowanceType(allowance.AllowanceType); err != nil {
			return err
		}
		if err := IsValidAllowanceValue(allowance.AllowanceType, allowance.Amount, *taxableIncome); err != nil {
			return err
		}

		deductionAmount := allowance.Amount
		if allowance.AllowanceType == models.DonationType && deductionAmount > 100000 {
			deductionAmount = 100000
		} else if allowance.AllowanceType == models.KReceiptType && deductionAmount > 50000 {
			deductionAmount = 50000
		}

		*taxableIncome -= deductionAmount
	}
	return nil
}

func IsValidAllowanceType(t string) error {
	switch t {
	case models.DonationType, models.WhtType, models.KReceiptType:
		return nil
	default:
		return errs.NewValidationError("invalid allowance type")
	}
}

func IsValidAllowanceValue(allowanceType string, value, income float64) error {
	fmt.Println("value: ", value, "income: ", income, "allowanceType: ", allowanceType)
	var errMsg string
	switch allowanceType {
	case models.KReceiptType:
		if value < 0 {
			errMsg = "k-receipt value must be greater than 0"
		}
	case models.DonationType:
		if value < 0 {
			errMsg = "donation value must be greater than 0"
		}
	case models.WhtType:
		if value < 0 {
			errMsg = "wht value must be greater than 0"
		}
		if value > income {
			errMsg = "wht value must be less than income"
		}
	default:
		errMsg = "invalid allowance type"
	}

	if errMsg != "" {
		return errs.NewValidationError(fmt.Sprintf("invalid allowance %s: %f", errMsg, value))
	}
	return nil
}
