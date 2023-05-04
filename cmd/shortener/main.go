package main

import (
	"github.com/IKostarev/yandex-go-dev/internal/config"
	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"github.com/IKostarev/yandex-go-dev/internal/router"
	"github.com/IKostarev/yandex-go-dev/internal/storage/fs"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Can't read config: ", err)
	}

	storage := fs.Load(cfg)
	app := handlers.App{Config: cfg, Storage: storage}

	r := router.NewRouter(app)

	log.Fatal(http.ListenAndServe(cfg.ServerAddr, r))
}
