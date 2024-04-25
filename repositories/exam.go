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
