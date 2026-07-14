package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func respond(w http.ResponseWriter, r *http.Request, status int, v any) {
	render.Status(r, status)
	render.JSON(w, r, v)
}

func respondError(w http.ResponseWriter, r *http.Request, status int, msg string) {
	render.Status(r, status)
	render.JSON(w, r, ErrorResponse{Error: msg})
}
