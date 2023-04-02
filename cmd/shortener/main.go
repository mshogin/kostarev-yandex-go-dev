package main

import (
	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.CompressHandler)

	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		panic(err)
	}
}
