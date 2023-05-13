package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"github.com/IKostarev/yandex-go-dev/internal/storage/database/postgres"
	"net/http"
)

func (a *App) PingHandler(w http.ResponseWriter, _ *http.Request) {
	if _, err := postgres.NewDB(a.Config.DatabaseDSN); err != nil {
		logger.Errorf("error connect to db: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
