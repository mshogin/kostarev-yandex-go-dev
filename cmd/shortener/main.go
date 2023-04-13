package main

import (
	"flag"
	"github.com/IKostarev/yandex-go-dev/internal/config"
	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	flag.Parse()

	cfg := config.LoadConfig()
	app := handlers.App{Config: cfg}

	r := chi.NewRouter()

	r.Get("/{id}", app.GetURLHandler)
	r.Post("/", app.CompressHandler)

	log.Fatal(http.ListenAndServe(cfg.Port, r))
}
