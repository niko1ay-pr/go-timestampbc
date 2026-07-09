package domain

import "time"

type BallotStatus string

const (
	BallotStatusPending   BallotStatus = "pending"
	BallotStatusConfirmed BallotStatus = "confirmed"
	BallotStatusInvalid   BallotStatus = "invalid"
)

type Ballot struct {
	ID            string       `json:"id" db:"id"`
	PollID        string       `json:"poll_id" db:"poll_id"`
	FlatID        string       `json:"flat_id" db:"flat_id"`
	RandomNum     int64        `json:"random_num" db:"random_num"`
	Hash          string       `json:"hash" db:"hash"` // SHA256(poll_id + answers + random_num)
	Answers       string       `json:"answers" db:"answers"`
	Status        BallotStatus `json:"status" db:"status"`
	CreatedAt     time.Time    `json:"created_at" db:"created_at"`
	ConfirmedAt   *time.Time   `json:"confirmed_at,omitempty" db:"confirmed_at"`
	BlockHash     *string      `json:"block_hash,omitempty" db:"block_hash"`
	TransactionID *string      `json:"transaction_id,omitempty" db:"transaction_id"`
}

type Proof struct {
	PollID    string `json:"poll_id"`
	BallotID  string `json:"ballot_id"`
	FlatID    string `json:"flat_id"`
	RandomNum string `json:"random_num"`
	Answers   string `json:"answers"`
}
