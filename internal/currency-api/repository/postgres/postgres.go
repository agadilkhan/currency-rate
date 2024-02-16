package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	var dest []entity.Currency
	query := `
				SELECT id, code, rate, updated_at 
				FROM currency`

	err := r.Db.SelectContext(ctx, &dest, query)
	if err != nil {
		return nil, fmt.Errorf("failed to SelectContext err: %v", err)
	}

	return &dest, nil
}

func (r *Repository) GetByCode(ctx context.Context, code string) (*entity.Currency, error) {
	var dest entity.Currency
	query := `
				SELECT id, code, rate, updated_at
				FROM currency 
				WHERE code=$1`

	args := []any{code}

	if err := r.Db.GetContext(ctx, &dest, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("not found err: %v", err)
		}
		return nil, fmt.Errorf("failed to GetContext err: %v", err)
	}

	return &dest, nil
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
