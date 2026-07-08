package sqlite

import (
	"context"
	"database/sql"
	"go-timestampbc/internal/models"
	"go-timestampbc/internal/store"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetByID(ctx context.Context, id string) (*models.Poll, error) {
	return nil, store.ErrNotFound
}
