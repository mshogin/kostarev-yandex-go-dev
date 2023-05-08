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
		_ = logger.Errorf("body is nil or empty: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	short, err := a.Storage.Save(string(body))
	if err != nil {
		_ = logger.Errorf("storage save is error: ", err)
		w.WriteHeader(http.StatusBadRequest) //TODO в будущем переделать на http.StatusInternalServerError
		return
	}

	long, err := url.JoinPath(a.Config.BaseShortURL, short)
	if err != nil {
		_ = logger.Errorf("join path have err: ", err)
		w.WriteHeader(http.StatusBadRequest) //TODO в будущем переделать на http.StatusInternalServerError
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(long))
	if err != nil {
		_ = logger.Errorf("Failed to send URL: ", err)
	}
}
