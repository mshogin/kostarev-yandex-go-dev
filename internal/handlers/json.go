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
		_ = logger.Errorf("json decode is error: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	short, err := a.Storage.Save(req.ServerURL)
	if err != nil {
		_ = logger.Errorf("storage save is error: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp.BaseShortURL, err = url.JoinPath(a.Config.BaseShortURL, short)
	if err != nil {
		_ = logger.Errorf("join path have err: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respContent, err := json.Marshal(resp)
	if err != nil {
		_ = logger.Errorf("json marshal is error: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(respContent); err != nil {
		_ = logger.Errorf("Failed to send URL on json handler: %w", err)
	}
}
