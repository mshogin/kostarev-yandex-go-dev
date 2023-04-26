package handlers

import (
	"compress/gzip"
	"github.com/IKostarev/yandex-go-dev/internal/storage"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strings"
)

func (a *App) GetURLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	url := chi.URLParam(r, "id")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m, err := storage.GetURL(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //TODO в будущем переделать на http.StatusNotFound
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
		if _, err := gzWriter.Write([]byte(m)); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		return
	}

	w.Header().Add("Location", m)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
