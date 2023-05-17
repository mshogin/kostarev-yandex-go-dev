package handlers

import (
	"encoding/json"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"io"
	"net/http"
	"net/url"
)

type URLsRequest struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type URLsResponse struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

func (a *App) BatchHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		logger.Errorf("body is nil or empty: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req URLsRequest
	var resp URLsResponse

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Errorf("json decode is error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	short, err := a.Storage.Save(req.OriginalURL, req.CorrelationID)
	if err != nil {
		logger.Errorf("batch save is error: %s", err)
		w.WriteHeader(http.StatusBadRequest) //TODO в будущем переделать на http.StatusInternalServerError
		return
	}

	resp.ShortURL, err = url.JoinPath(a.Config.BaseShortURL, short)
	if err != nil {
		logger.Errorf("join path have err: %s", err)
		w.WriteHeader(http.StatusBadRequest) //TODO в будущем переделать на http.StatusInternalServerError
		return
	}

	_, resp.CorrelationID = a.Storage.Get(resp.ShortURL, req.CorrelationID)

	respContent, err := json.Marshal(resp)
	if err != nil {
		logger.Errorf("json marshal is error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(respContent); err != nil {
		logger.Errorf("Failed to send URLsResponse on batch handler: %s", err)
	}
}
