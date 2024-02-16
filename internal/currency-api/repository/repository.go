package repository

import (
	"context"
	"github.com/agadilkhan/currency-rate/internal/currency-api/entity"
)

type Repository interface {
	List(ctx context.Context) (*[]entity.Currency, error)
	GetByCode(ctx context.Context, code string) (*entity.Currency, error)
	Update(ctx context.Context, currency *entity.Currency) error
}
