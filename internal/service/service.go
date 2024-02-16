package service

import (
	"context"
	"fmt"
	"github.com/agadilkhan/currency-rate/internal/entity"
	"github.com/agadilkhan/currency-rate/internal/repository"
	"github.com/agadilkhan/currency-rate/internal/transport"
)

type Service struct {
	transport *transport.HttpTransport
	repo      repository.Repository
}

func New(repo repository.Repository, httpTransport *transport.HttpTransport) *Service {
	return &Service{
		repo:      repo,
		transport: httpTransport,
	}
}

func (s *Service) List(ctx context.Context) (*[]entity.Currency, error) {
	return s.repo.List(ctx)
}

func (s *Service) GetByCode(ctx context.Context, code string) (*entity.Currency, error) {
	return s.repo.GetByCode(ctx, code)
}

func (s *Service) Update(ctx context.Context) error {
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

func (s *Service) Save(ctx context.Context) error {
	rates, err := s.transport.GetCurrencies()
	if err != nil {
		return fmt.Errorf("failed to GetCurrencies err: %v", err)
	}

	for _, currency := range *rates {
		_, err = s.repo.Save(ctx, &currency)
		if err != nil {
			return fmt.Errorf("failed to Save err: %v", err)
		}
	}

	return nil
}
