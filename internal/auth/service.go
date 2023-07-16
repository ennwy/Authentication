package auth

import (
	"context"
	"fmt"

	"github.com/ennwy/auth/internal/app"
)

type Service struct {
	ctx     context.Context
	Storage app.Storage
}

var l app.Logger

func NewService(ctx context.Context, log app.Logger, storage app.Storage) (*Service, error) {
	l = log

	s := &Service{
		Storage: storage,
	}

	if err := s.Storage.Connect(ctx); err != nil {
		return nil, fmt.Errorf("storage conn: %w", err)
	}

	return s, nil
}
