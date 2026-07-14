package handler

import (
	"context"
	"encoding/json"
	"errors"
	"go-timestampbc/internal/domain"
	"go-timestampbc/internal/store"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type PollGetter interface {
	GetByID(ctx context.Context, id string) (*domain.Poll, error)
}

func HandleGetPoll(pollGetter PollGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pollID := r.PathValue("pollId")
		if pollID == "" {
			respond(w, r, http.StatusInternalServerError, "pollId is required")
			return
		}

		poll, err := pollGetter.GetByID(r.Context(), pollID)
		if err != nil {
			if errors.Is(err, store.ErrNotFound) {
				respondError(w, r, http.StatusNotFound, "poll not found")
				return
			}
			respondError(w, r, http.StatusInternalServerError, "internal server error")
			return
		}

		respond(w, r, http.StatusOK, poll)
	}
}

type PollCreater interface {
	Create(ctx context.Context, poll *domain.Poll) error
}

func HandleCreatePoll(sl *slog.Logger, pollCreater PollCreater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		poll := &domain.Poll{}
		json.NewDecoder(r.Body).Decode(poll)
		poll.ID = uuid.New().String()
		poll.CreatedAt = time.Now().UTC().Truncate(time.Second)
		err := pollCreater.Create(r.Context(), poll)
		if err != nil {
			respondError(w, r, http.StatusInternalServerError, "creating error")
			sl.Error("creating error", slog.String("error", err.Error()))
			return
		}

		respond(w, r, http.StatusCreated, poll)
	}
}
