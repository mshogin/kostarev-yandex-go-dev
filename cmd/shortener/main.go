package main

import (
	"github.com/IKostarev/yandex-go-dev/config"
	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	r := chi.NewRouter()

	r.Get("/{id}", handlers.GetURLHandler)
	r.Post("/", handlers.CompressHandler)

	err := http.ListenAndServe(cfg.HTTPAddr, r)
	if err != nil {
		log.Fatal(err)
	}
}
