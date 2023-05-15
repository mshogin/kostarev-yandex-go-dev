package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"github.com/IKostarev/yandex-go-dev/internal/storage/database/postgres"
	"net/http"
)

func (a *App) PingHandler(w http.ResponseWriter, _ *http.Request) {
	db, err := postgres.NewPostgresDB(a.Config.FileStoragePath)
	if err != nil {
		logger.Errorf("error ping handler: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if db.Pool() {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
