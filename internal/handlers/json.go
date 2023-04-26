package handlers

import (
	"compress/gzip"
	"encoding/json"
	"github.com/IKostarev/yandex-go-dev/internal/storage"
	"log"
	"net/http"
	url1 "net/url"
	"strings"
)

type URL struct {
	ServerURL string `json:"url"`
}

type Result struct {
	BaseShortURL string `json:"result"`
}

func (a *App) JSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var url URL
	var res Result

	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	miniURL := storage.SaveURL(url.ServerURL)

	var err error
	res.BaseShortURL, err = url1.JoinPath(a.Config.BaseShortURL, miniURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(map[string]string{"result": res.BaseShortURL})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")

		gz := gzip.NewWriter(w)
		defer func() {
			if err := gz.Close(); err != nil {
				log.Println("Error closing gzip writer:", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}()

		gzWriter := &GzipResponseWriter{Writer: gz, ResponseWriter: w}
		if _, err := gzWriter.Write(resp); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		w.WriteHeader(http.StatusCreated)
		if _, err := w.Write(resp); err != nil {
			log.Fatal("Failed to send URL on json handler")
		}
	}
}
