package service

import (
	"context"
	"github.com/agadilkhan/currency-rate/internal/currency-api/entity"
)

type UseCase interface {
	List(ctx context.Context) (*[]entity.Currency, error)
	GetByCode(ctx context.Context, code string) (*entity.Currency, error)
	ForceUpdate(ctx context.Context) error
}
