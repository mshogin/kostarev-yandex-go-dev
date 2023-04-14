package main

import (
	"github.com/IKostarev/yandex-go-dev/internal/config"
	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {

	cfg := config.LoadConfig()
	app := handlers.App{Config: cfg}

	r := chi.NewRouter()

	r.Get("/{id}", app.GetURLHandler)
	r.Post("/", app.CompressHandler)

	log.Printf("server starting on port: %v", *cfg.ServerAddr)

	log.Fatal(http.ListenAndServe(*cfg.ServerAddr, r))
}
