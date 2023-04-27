package middleware

import (
	"github.com/IKostarev/yandex-go-dev/internal/handlers"
	"log"
	"net/http"
	"strings"
)

func GzipMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		originWriter := w
		acceptEncoding := r.Header.Get("Accept-Encoding")
		contentEncoding := r.Header.Get("Content-Encoding")
		suppAccept := strings.Contains(acceptEncoding, "gzip")
		suppContent := strings.Contains(contentEncoding, "gzip")

		if suppContent && suppAccept {
			compressWriter := handlers.NewCompressWriter(w)
			originWriter = compressWriter
			defer func() {
				if err := compressWriter.Close(); err != nil {
					log.Println("Error compress write:", err)
					return
				}
			}()

			compressReader, err := handlers.NewCompressReader(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest) //TODO поменять на 500
			}

			r.Body = compressReader
			defer func() {
				if err := compressReader.Close(); err != nil {
					log.Println("Error compress reade:", err)
					return
				}
			}()
		}
		h.ServeHTTP(originWriter, r)
	}
}
