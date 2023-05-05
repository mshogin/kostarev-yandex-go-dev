package handlers

import (
	"fmt"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *App) GetURLHandler(w http.ResponseWriter, r *http.Request) {
	url := chi.URLParam(r, "id")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("url param bad with id: ", url)
		return
	}

	m := a.Storage.Get(url)

	fmt.Println("url geturl = ", url)
	fmt.Println("m geturl = ", m)

	if m == "" {
		w.WriteHeader(http.StatusBadRequest) //TODO в будущем переделать на http.StatusNotFound
		logger.Error("get url is bad: ", m)
		return
	}

	w.Header().Add("Location", m)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
