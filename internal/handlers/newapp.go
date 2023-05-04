package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/config"
	"github.com/IKostarev/yandex-go-dev/internal/storage"
)

type App struct {
	Config  config.Config
	Storage storage.Storage
}
