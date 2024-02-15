package service

import (
	"context"
	"github.com/agadilkhan/currency-rate/internal/currency-api/entity"
	"github.com/agadilkhan/currency-rate/internal/currency-api/repository"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List(ctx context.Context) (*[]entity.Currency, error) {
	return nil, nil
}

func (s *Service) GetByCode(ctx context.Context, code string) (*entity.Currency, error) {
	return nil, nil
}

func (s *Service) ForceUpdate(ctx context.Context) error {
	return nil
}
