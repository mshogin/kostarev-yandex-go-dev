package main

import (
	"github.com/IKostarev/yandex-go-dev/internal/config"
	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"github.com/IKostarev/yandex-go-dev/internal/router"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Can't read config: %w", err)
	}

	app := handlers.App{Config: cfg}

	r := router.NewRouter(app)

	log.Fatal(http.ListenAndServe(cfg.ServerAddr, r))
}
