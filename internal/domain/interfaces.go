package domain

import (
	"context"
)

type Storage interface {
	Polls() PollStorage
	Ballots() BallotStorage
}

type PollStorage interface {
	GetByID(ctx context.Context, id string) (*Poll, error)
}

type BallotStorage interface {
}
