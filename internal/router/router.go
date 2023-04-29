package router

import (
	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"github.com/IKostarev/yandex-go-dev/internal/middleware/gzip"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(app handlers.App) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(gzip.Request)
	r.Use(gzip.Response)

	r.Route("/", func(r chi.Router) {
		r.Post("/", app.CompressHandler)
		r.Get("/{id}", app.GetURLHandler)

		r.Route("/api", func(r chi.Router) {
			r.Post("/shorten", app.JSONHandler)
		})
	})

	return r
}
