package service

import (
	"context"
	"github.com/agadilkhan/currency-rate/internal/currency-api/entity"
)

type UseCase interface {
	List(ctx context.Context) (*[]entity.Currency, error)
	GetByCode(ctx context.Context, code string) (*entity.Currency, error)
	Update(ctx context.Context) error
	Save(ctx context.Context, currency *entity.Currency) error
}
