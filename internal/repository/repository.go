package repository

import (
	"context"
	"github.com/agadilkhan/currency-rate/internal/entity"
)

type Repository interface {
	List(ctx context.Context) (*[]entity.Currency, error)
	GetByCode(ctx context.Context, code string) (*entity.Currency, error)
	Update(ctx context.Context, currency *entity.Currency) error
	Save(ctx context.Context, currency *entity.Currency) (int, error)
}
