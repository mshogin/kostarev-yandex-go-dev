package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/storage"
	"net/http"
)

func (a *App) GetURLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	url := r.URL.String()
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	m, err := storage.GetURL(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Add("Location", m)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
