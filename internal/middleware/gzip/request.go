package gzip

import (
	"compress/gzip"
	"github.com/IKostarev/yandex-go-dev/internal/logger"
	"net/http"
	"strings"
)

func Request(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				logger.Error("new reader is error: ", err)
				return
			}

			r.Body = gz
			defer func() {
				if err = gz.Close(); err != nil {
					logger.Error("GzipRequest gz.Close() failed: %v", err)
				}
			}()
		}

		next.ServeHTTP(w, r)
	})
}
