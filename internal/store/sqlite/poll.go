package sqlite

import (
	"context"
	"database/sql"
	"go-timestampbc/internal/domain"
	"go-timestampbc/internal/store"
)

type PollStore struct {
	db *sql.DB
}

func NewPollStore(db *sql.DB) *PollStore {
	return &PollStore{db: db}
}

func (p *PollStore) GetByID(ctx context.Context, id string) (*domain.Poll, error) {
	return nil, store.ErrNotFound
}
