package main

import (
	"github.com/IKostarev/yandex-go-dev/internal/config"
	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Can't read config: %w", err)
	}

	app := handlers.App{Config: cfg}

	r := chi.NewRouter()

	r.Get("/{id}", logger.ResponseLogger(app.GetURLHandler))
	r.Post("/", logger.RequestLogger(app.CompressHandler))

	log.Fatal(http.ListenAndServe(cfg.ServerAddr, r))
}
