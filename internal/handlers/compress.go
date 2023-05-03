package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"github.com/IKostarev/yandex-go-dev/internal/service"
	"io"
	"net/http"
	"net/url"
)

func (a *App) CompressHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("body is nil or empty: ", err)
		return
	}

	miniURL := service.SaveURL(string(body))

	newURL, err := url.JoinPath(a.Config.BaseShortURL, miniURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //TODO в будущем переделать на http.StatusInternalServerError
		logger.Error("join path have err: ", err)
		return
	}

	a.FileStorage(miniURL, string(body))

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(newURL))
	if err != nil {
		logger.Error("Failed to send URL: ", err)
	}
}
