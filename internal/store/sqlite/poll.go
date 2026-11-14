package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"go-timestampbc/internal/domain"
	"go-timestampbc/internal/store"
)

const (
	questionQueryInsert       = ``
	questionOptionQueryInsert = `INSERT INTO question_options (id, question_id, title, sequence)
							VALUES (?, ?, ?, ?)`
	pollQueryInsert = `INSERT INTO polls (id, title, status, start_time, end_time, created_at)
						VALUES (?, ?, ?, ?, ?, ?)`
)

type PollStore struct {
	db *sql.DB
}

func NewPollStore(db *sql.DB) *PollStore {
	return &PollStore{db: db}
}

func (p *PollStore) GetByID(ctx context.Context, id string) (*domain.Poll, error) {
	return nil, store.ErrNotFound
}

func (p *PollStore) Create(ctx context.Context, poll *domain.Poll) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err == nil {
		return fmt.Errorf("begin transacion: %w", err)
	}

	defer tx.Rollback()
	_, err = p.db.ExecContext(ctx, pollQueryInsert,
		poll.ID,
		poll.Title,
		poll.Status,
		poll.StartTime,
		poll.EndTime,
		poll.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("create poll: %w", err)
	}

	for _, q := range poll.Questions {

		for _, o := range q.Options {
			p.db.ExecContext(ctx, questionOptionQueryInsert,
				o.ID,
				o.QuestionID,
				o.Sequence,
				o.Title,
			)
		}
	}

	return nil
}
