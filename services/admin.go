package services

import (
	"context"

	"github.com/imrinry/assessment-tax/errs"
	"github.com/imrinry/assessment-tax/logs"
)

func (s *service) DeductionPersonalSetting(ctx context.Context, amount float64) (interface{}, error) {
	const keyName = "personalDeduction"
	if amount < 0 {
		return nil, errs.NewValidationError("personalDeduction must be greater than 0")
	}
	if amount > 100000 {
		return nil, errs.NewValidationError("personalDeduction must be less than  or equal to 100000")
	}
	_, err := s.r.UpdateDeductionSetting(ctx, keyName, amount)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	return s.r.UpdateDeductionSetting(ctx, keyName, amount)
}
