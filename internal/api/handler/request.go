package handler

import (
	"go-timestampbc/internal/domain"
	"time"
)

type PollCreateRequest struct {
	Title     string            `json:"title"`
	Status    domain.PollStatus `json:"status"`
	StartTime time.Time         `json:"start_time"`
	EndTime   time.Time         `json:"end_time"`
	CreatedAt time.Time         `json:"created_at"`
	Questions *domain.Question  `json:"questions"`
}
