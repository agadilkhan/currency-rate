package service

import (
	"context"
)

type UseCase interface {
	Save(ctx context.Context) error
	Update(ctx context.Context) error
}
