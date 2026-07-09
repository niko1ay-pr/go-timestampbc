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
	Create(ctx context.Context, poll *Poll) error
}

type BallotStorage interface {
}
