package main

import (
	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", handlers.GetUrlHandler)
	r.Post("/{id}", handlers.CompressHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
