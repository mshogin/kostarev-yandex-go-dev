package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/storage/database/postgres"
	"net/http"
)

func (a *App) PingHandler(w http.ResponseWriter, _ *http.Request) {
	db := &postgres.DB{}
	if db.Pool() {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
