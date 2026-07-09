package handler

import (
	"context"
	"go-timestampbc/internal/domain"
	"net/http"
)

type BallotGetter interface {
	GetByID(ctx context.Context, id string) (*domain.Ballot, error)
}

func HandleCreateBallot(ballotGetter BallotGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := "1" // TODO
		ballotGetter.GetByID(r.Context(), id)
	}
}
