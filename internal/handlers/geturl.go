package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *App) GetURLHandler(w http.ResponseWriter, r *http.Request) {
	url := chi.URLParam(r, "id")
	if url == "" {
		_ = logger.Errorf("url param bad with id: ", url)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := a.Storage.Get(url)

	if m == "" {
		_ = logger.Errorf("get url is bad: ", m)
		w.WriteHeader(http.StatusBadRequest) //TODO в будущем переделать на http.StatusNotFound
		return
	}

	w.Header().Add("Location", m)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
