package main

import (
	"github.com/IKostarev/yandex-go-dev/config"
	"log"
	"net/http"

	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.LoadConfig()

	r := chi.NewRouter()

	r.Get("/{id}", handlers.GetURLHandler)
	r.Post("/", handlers.CompressHandler)

	log.Fatal(http.ListenAndServe(cfg.HTTPAddr, r))
}
