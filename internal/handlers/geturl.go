package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/storage"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *App) GetURLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	url := chi.URLParam(r, "id")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	m, err := storage.GetURL(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //TODO в будущем переделать на http.StatusNotFound
	}

	w.Header().Add("Location", m)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
