package services

import (
	"context"

	"github.com/imrinry/assessment-tax/repositories"
)

type Services interface {
	ExamServices(ctx context.Context) (interface{}, error)
}

type service struct {
	r repositories.Repositories
}

func New(r repositories.Repositories) Services {
	return &service{
		r: r,
	}
}
