package sqlite

import (
	"database/sql"
	"go-timestampbc/internal/domain"
)

type Store struct {
	polls   *PollStore
	ballots *BallotStore
}

func NewStore(db *sql.DB) *Store {
	return &Store{polls: NewPollStore(db), ballots: NewBallotStore(db)}
}

func (s *Store) Polls() domain.PollStorage {
	return s.polls
}

func (s *Store) Ballots() domain.BallotStorage {
	return s.ballots
}
