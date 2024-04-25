package services

import (
	"context"

	"github.com/imrinry/assessment-tax/models"
	"github.com/imrinry/assessment-tax/repositories"
)

type Services interface {
	TaxCalculations(ctx context.Context, income float64, wht float64, allowances []models.Allowance) (models.TaxResponse, error)
	DeductionPersonalSetting(ctx context.Context, amount float64) (interface{}, error)
}

type service struct {
	r repositories.Repositories
}

func New(r repositories.Repositories) Services {
	return &service{
		r: r,
	}
}
