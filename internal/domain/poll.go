package domain

import "time"

type PollStatus string

const (
	PollStatusUpcoming PollStatus = "upcoming"
	PollStatusActive   PollStatus = "active"
	PollStatusEnded    PollStatus = "ended"
)

type Poll struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Status    PollStatus `json:"status"`
	StartTime time.Time  `json:"start_time"`
	EndTime   time.Time  `json:"end_time"`
	CreatedAt time.Time  `json:"created_at"`
}
