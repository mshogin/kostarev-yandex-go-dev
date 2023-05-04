package handlers

import (
	"encoding/json"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"net/http"
	"net/url"
)

type URLRequest struct {
	ServerURL string `json:"url"`
}

type ResultResponse struct {
	BaseShortURL string `json:"result"`
}

func (a *App) JSONHandler(w http.ResponseWriter, r *http.Request) {
	var req URLRequest
	var resp ResultResponse

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("json decode is error: ", err)
		return
	}

	short, err := a.Storage.Save(req.ServerURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("storage save is error: ", err)
		return
	}

	resp.BaseShortURL, err = url.JoinPath(a.Config.BaseShortURL, short)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("join path have err: ", err)
		return
	}

	respContent, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error("json marshal is error: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(respContent); err != nil {
		logger.Error("Failed to send URL on json handler", err)
	}
}
