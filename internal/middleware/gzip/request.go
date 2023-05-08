package gzip

import (
	"compress/gzip"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"net/http"
	"strings"
)

// TODO тест
func Request(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				_ = logger.Errorf("new reader is error: %w", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			r.Body = gz
			defer func() {
				if err = gz.Close(); err != nil {
					_ = logger.Errorf("GzipRequest gz.Close() failed: %w", err)
				}
			}()
		}

		next.ServeHTTP(w, r)
	})
}
