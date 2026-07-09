package handler

import (
	"context"
	"encoding/json"
	"errors"
	"go-timestampbc/internal/domain"
	"go-timestampbc/internal/store"
	"net/http"
)

type PollGetter interface {
	GetByID(ctx context.Context, id string) (*domain.Poll, error)
}

func HandleGetPoll(pollGetter PollGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pollID := r.PathValue("pollId")
		if pollID == "" {
			writeJSON(w, http.StatusInternalServerError, apiError("pollId is required"))
			return
		}

		poll, err := pollGetter.GetByID(r.Context(), pollID)
		if err != nil {
			if errors.Is(err, store.ErrNotFound) {
				writeJSON(w, http.StatusNotFound, apiError("poll not found"))
				return
			}
			writeJSON(w, http.StatusInternalServerError, apiError("internal server error"))
			return
		}

		writeJSON(w, http.StatusOK, poll)
	}
}

type PollCreater interface {
	Create(ctx context.Context, poll *domain.Poll) error
}

func HandleCreatePoll(pollCreater PollCreater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// read poll from json
		poll := &domain.Poll{}
		err := pollCreater.Create(r.Context(), poll)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, apiError("creating error"))
			return
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func apiError(s string) *ErrorResponse {
	return &ErrorResponse{Error: s}
}
