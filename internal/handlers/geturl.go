package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/service"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *App) GetURLHandler(w http.ResponseWriter, r *http.Request) {
	url := chi.URLParam(r, "id")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m, err := service.GetURL(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //TODO в будущем переделать на http.StatusNotFound
		return
	}

	w.Header().Add("Location", m)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
