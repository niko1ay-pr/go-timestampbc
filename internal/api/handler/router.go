package handler

import (
	mw "go-timestampbc/internal/api/middleware"
	"go-timestampbc/internal/store"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(s store.Store, sl *slog.Logger) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(mw.Logger(sl))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// GET /polls/{pollId}
	r.Get("/polls/{pollId}", HandleGetPoll(s))
	// Health-check endpoint
	r.Get("/health", func(w http.ResponseWriter, h *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok", "timestamp":"` + time.Now().Format(time.RFC3339) + `"}`))
	})

	return r
}
