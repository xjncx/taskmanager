package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(handler *Handler) http.Handler {

	r := chi.NewRouter()

	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", handler.CreateTask)
		r.Get("/{id}", handler.GetTask)
		r.Delete("/{id}", handler.DeleteTask)
	})

	return r
}
