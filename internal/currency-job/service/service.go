package service

import (
	"context"
	"fmt"
	"github.com/agadilkhan/currency-rate/internal/currency-job/repository"
	"github.com/agadilkhan/currency-rate/internal/currency-job/transport"
)

type Service struct {
	transport transport.HttpTransport
	repo      repository.Repository
}

func New(repo repository.Repository, transport transport.HttpTransport) *Service {
	return &Service{
		repo:      repo,
		transport: transport,
	}
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
