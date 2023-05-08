package main

import (
	"github.com/IKostarev/yandex-go-dev/internal/config"
	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	store "github.com/IKostarev/yandex-go-dev/internal/storage"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatalf("Can't read config: ", err)
	}

	storage, err := store.NewStorage(cfg)
	if err != nil {
		logger.Fatalf("Can't storage download", err)
	}

	app := handlers.NewApp(cfg, storage)
	log.Fatal(http.ListenAndServe(cfg.ServerAddr, app))
}
