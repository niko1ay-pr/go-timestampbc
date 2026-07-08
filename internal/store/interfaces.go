package store

import (
	"context"
	"go-timestampbc/internal/models"
)

type PollStore interface {
	GetByID(ctx context.Context, id string) (*models.Poll, error)
}

type Store interface {
	PollStore
}
