package handlers

import (
	"encoding/json"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"github.com/IKostarev/yandex-go-dev/internal/service"
	"net/http"
	url1 "net/url"
)

type URL struct {
	ServerURL string `json:"url"`
}

type Result struct {
	BaseShortURL string `json:"result"`
}

func (a *App) JSONHandler(w http.ResponseWriter, r *http.Request) {
	var url URL
	var res Result

	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("json decode is error: ", err)
		return
	}

	miniURL := service.SaveURL(url.ServerURL)

	var err error
	res.BaseShortURL, err = url1.JoinPath(a.Config.BaseShortURL, miniURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("join path have err: ", err)
		return
	}

	resp, err := json.Marshal(map[string]string{"result": res.BaseShortURL})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("json marshal is error: ", err)
		return
	}

	a.FileStorage(miniURL, url.ServerURL)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(resp); err != nil {
		logger.Error("Failed to send URL on json handler", err)
	}
}
