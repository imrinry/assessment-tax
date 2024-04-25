package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Repositories interface {
	ExamRepo(ctx context.Context) (interface{}, error)
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repositories {
	return &repository{
		db: db,
	}
}
