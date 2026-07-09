package sqlite

import (
	"database/sql"
)

type BallotStore struct {
	db *sql.DB
}

func NewBallotStore(db *sql.DB) *BallotStore {
	return &BallotStore{db: db}
}
