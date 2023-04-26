package handlers

import (
	"compress/gzip"
	"github.com/IKostarev/yandex-go-dev/internal/storage"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (a *App) CompressHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer func() {
			if err := gz.Close(); err != nil {
				log.Println("Error closing gzip writer:", err)
				return
			}
		}()

		w = &GzipResponseWriter{Writer: gz, ResponseWriter: w}
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	miniURL := storage.SaveURL(string(body))

	newURL, err := url.JoinPath(a.Config.BaseShortURL, miniURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //TODO в будущем переделать на http.StatusInternalServerError
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := io.WriteString(w, newURL); err != nil {
		log.Fatal("Failed to send URL")
	}
}
