package models

import "time"

type VoteStatus string

const (
	VoteStatusPending   VoteStatus = "pending"
	VoteStatusConfirmed VoteStatus = "confirmed"
	VoteStatusInvalid   VoteStatus = "invalid"
)

type Vote struct {
	ID            uint64     `json:"id" db:"id"`
	Election_id   uint64     `json:"election_id" db:"election_id"`
	Candidate_id  uint64     `json:"candidate_id" db:"candidate_id"`
	Random_num    int64      `json:"random_num" db:"random_num"`
	Hash          string     `json:"hash" db:"hash"`
	Status        string     `json:"status" db:"status"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	ConfirmedAt   *time.Time `json:"confirmed_at,omitempty" db:"confirmed_at"`
	BlockHash     *string    `json:"block_hash,omitempty" db:"block_hash"`
	TransactionID *string    `json:"transaction_id,omitempty" db:"transaction_id"`
}
