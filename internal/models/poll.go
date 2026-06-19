package models

import "time"

type PollStatus string

const (
	PollStatusUpcoming PollStatus = "upcoming"
	PollStatusActive   PollStatus = "active"
	PollStatusEnded    PollStatus = "ended"
)

type Poll struct {
	ID        string     `json:"id" db:"id"`
	Title     string     `json:"title" db:"title"`
	Status    PollStatus `json:"status" db:"status"`
	StartTime time.Time  `json:"start_time" db:"start_time"`
	EndTime   time.Time  `json:"end_time" db:"end_time"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
}
