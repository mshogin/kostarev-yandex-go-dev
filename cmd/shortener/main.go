package main

import (
	"fmt"
	"github.com/IKostarev/yandex-go-dev/config"
	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Println(cfg)

	r := chi.NewRouter()

	r.Get("/{id}", handlers.GetURLHandler)
	r.Post("/", handlers.CompressHandler)

	err := http.ListenAndServe(cfg.Port, r)
	if err != nil {
		log.Fatal(err)
	}
}
