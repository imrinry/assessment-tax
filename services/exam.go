package services

import (
	"context"

	"github.com/imrinry/assessment-tax/errs"
	"github.com/imrinry/assessment-tax/logs"
)

func (s *service) ExamServices(ctx context.Context) (interface{}, error) {
	var err error
	data, err := s.r.ExamRepo(ctx)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	// number := 200
	// if number > 200 {
	// 	return nil, errors.New("number is gte 200")
	// }
	state := true
	if state {
		return nil, errs.NewNotFoundError("")
	}
	return data, nil
}
