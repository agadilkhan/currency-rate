package postgres

import (
	"context"
	"github.com/agadilkhan/currency-rate/internal/currency-api/entity"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db,
	}
}

func (r *Repository) List(ctx context.Context) (*[]entity.Currency, error) {
	return nil, nil
}

func (r *Repository) GetByCode(ctx context.Context, code string) (*entity.Currency, error) {
	return nil, nil
}

func (r *Repository) Update(ctx context.Context) error {
	return nil
}
