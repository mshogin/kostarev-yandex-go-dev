package main

import (
	"github.com/IKostarev/yandex-go-dev/internal/config"
	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"github.com/IKostarev/yandex-go-dev/internal/router"
	store "github.com/IKostarev/yandex-go-dev/internal/storage"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Can't read config: ", err)
	}

	storage, err := store.NewStorage(cfg)
	if err != nil {
		logger.Error("Can't storage download", err)
	}

	app := handlers.App{Config: cfg, Storage: storage}

	r := router.NewRouter(app)

	log.Fatal(http.ListenAndServe(cfg.ServerAddr, r))
}
