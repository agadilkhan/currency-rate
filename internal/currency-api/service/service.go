package service

import (
	"context"
	"fmt"
	"github.com/agadilkhan/currency-rate/internal/currency-api/entity"
	"github.com/agadilkhan/currency-rate/internal/currency-api/repository"
	"github.com/agadilkhan/currency-rate/internal/currency-api/transport"
)

type Service struct {
	transport transport.HttpTransport
	repo      repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) List(ctx context.Context) (*[]entity.Currency, error) {
	return s.repo.List(ctx)
}

func (s *Service) GetByCode(ctx context.Context, code string) (*entity.Currency, error) {
	return s.repo.GetByCode(ctx, code)
}

func (s *Service) ForceUpdate(ctx context.Context) error {
	rates, err := s.transport.GetCurrencies()
	if err != nil {
		return fmt.Errorf("failed to GetCurrencies err: %v", err)
	}

	for _, currency := range *rates {
		err = s.repo.Update(ctx, &currency)
		if err != nil {
			return fmt.Errorf("failed to Update err: %v", err)
		}
	}
	return nil
}
