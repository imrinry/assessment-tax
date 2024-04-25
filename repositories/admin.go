package repositories

import (
	"context"
)

func (r *repository) ExamRepo(ctx context.Context) (interface{}, error) {
	var err error

	// ? open below code if want to return error
	// err = errors.New("test exam error")
	if err != nil {
		return nil, err
	}

	return "Exam Response", nil
}

func (r *repository) UpdateDeductionSetting(ctx context.Context, name string, amount float64) (interface{}, error) {
	_, err := r.db.Exec(`update deductions set value = $1  where  name = $2`, amount, name)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
