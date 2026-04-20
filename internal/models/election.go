package models

import "time"

type ElectionStatus string

const (
	ElectionStatusUpcoming ElectionStatus = "upcoming"
	ElectionStatusActive   ElectionStatus = "active"
	ElectionStatusEnded    ElectionStatus = "ended"
)

type Election struct {
	ID        string         `json:"id" db:"id"`
	Name      string         `json:"name" db:"name"`
	Status    ElectionStatus `json:"status" db:"status"`
	StartTime time.Time      `json:"start_time" db:"start_time"`
	EndTime   time.Time      `json:"end_time" db:"end_time"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
}
