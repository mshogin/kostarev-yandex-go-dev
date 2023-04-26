package middleware

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"io"
	"log"
	"net/http"
	"strings"
)

func GzipMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

			if strings.Contains(r.Header.Get("Content-Type"), "application/json") {

				var buf bytes.Buffer
				tee := io.TeeReader(r.Body, &buf)

				var data interface{}
				err := json.NewDecoder(tee).Decode(&data)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				r.Body = io.NopCloser(&buf)

				h(&handlers.GzipResponseWriter{Writer: gz, ResponseWriter: w}, r)

				if err = gz.Close(); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				w.Header().Del("Content-Encoding")
				w.Header().Set("Content-Type", "application/json")

				if _, err = w.Write(buf.Bytes()); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			} else {
				gzWriter := &handlers.GzipResponseWriter{Writer: gz, ResponseWriter: w}
				h.ServeHTTP(gzWriter, r)
			}
		} else {
			h.ServeHTTP(w, r)
		}
	}
}
