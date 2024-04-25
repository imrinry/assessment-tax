package repositories

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repositories interface {
	ExamRepo(ctx context.Context) (interface{}, error)
	UpdateDeductionSetting(ctx context.Context, name string, amount float64) (interface{}, error)
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repositories {
	return &repository{
		db: db,
	}
}
