package api

import "net/http"

func NewRouter(handler *Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", handler.CreateTask)
		r.Get("/{id}", handler.GetTask)
		r.Delete("/{id}", handler.DeleteTask)
	})

	return r
}
