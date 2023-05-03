package router

import (
	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"github.com/IKostarev/yandex-go-dev/internal/middleware/gzip"
	"github.com/go-chi/chi/v5"
)

func NewRouter(app handlers.App) chi.Router {
	r := chi.NewRouter()

	r.Use(gzip.Request)
	r.Use(gzip.Response)

	r.Route("/", func(r chi.Router) {
		r.Post("/", logger.RequestLogger(app.CompressHandler))
		r.Get("/{id}", logger.ResponseLogger(app.GetURLHandler))

		r.Route("/api", func(r chi.Router) {
			r.Post("/shorten", app.JSONHandler)
		})
	})

	return r
}
