package postgres

import (
	"context"
	"fmt"
	"github.com/agadilkhan/currency-rate/internal/currency-job/entity"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Db: db,
	}
}

func (r *Repository) Save(ctx context.Context, currency *entity.Currency) (int, error) {
	query := `
		INSERT INTO currency (code, rate)
		VALUES ($1, $2)
		RETURNING id`

	args := []any{currency.Code, currency.Rate}

	var id int

	err := r.Db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to QueryRowContext err: %v", err)
	}

	return id, nil
}

func (r *Repository) Update(ctx context.Context, currency *entity.Currency) error {
	query := `
				UPDATE currency
				SET rate=$1 
				WHERE code=$2`

	_, err := r.Db.Exec(query, currency.Rate, currency.Code)
	if err != nil {
		return fmt.Errorf("failed to Exec err: %v", err)
	}

	return nil
}
