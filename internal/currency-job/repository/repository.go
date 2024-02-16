package repository

import (
	"context"
	"github.com/agadilkhan/currency-rate/internal/currency-job/entity"
)

type Repository interface {
	Save(ctx context.Context, currency *entity.Currency) (int, error)
	Update(ctx context.Context, currency *entity.Currency) error
}
