package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/logger"
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

	short, err := a.Storage.Save(string(body))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //TODO в будущем переделать на http.StatusInternalServerError
		logger.Error("storage save is error: ", err)
		return
	}

	long, err := url.JoinPath(a.Config.BaseShortURL, short)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //TODO в будущем переделать на http.StatusInternalServerError
		logger.Error("join path have err: ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(long))
	if err != nil {
		logger.Error("Failed to send URL: ", err)
	}
}
