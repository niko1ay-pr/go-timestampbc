package sqlite

import (
	"context"
	"database/sql"
	"go-timestampbc/internal/domain"
)

type BallotStore struct {
	db *sql.DB
}

func NewBallotStore(db *sql.DB) *BallotStore {
	return &BallotStore{db: db}
}

type BallotFinder interface {
	GetByID(ctx context.Context, id string) (*domain.Ballot, error)
}
