package middleware

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"strings"
)

type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *gzipResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.Header().Del("Content-Length")
	w.ResponseWriter.WriteHeader(code)
}

func GzipMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			h.ServeHTTP(w, r)
			return
		}

		cntEncoding := r.Header.Get("Content-Encoding")
		if !strings.Contains(cntEncoding, "gzip") {
			h.ServeHTTP(w, r)
			return
		}

		cntType := r.Header.Get("Content-Type")
		jsonType := strings.Contains(cntType, "application/json")
		textType := strings.Contains(cntType, "text/plain")

		if jsonType || textType {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				log.Printf("Error creating gzip reader: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			defer func() {
				if err := gz.Close(); err != nil {
					log.Printf("Error closing gzip reader: %v", err)
					return
				}
			}()
			r.Body = http.MaxBytesReader(w, gz, r.ContentLength)
		}

		w.Header().Set("Content-Encoding", "gzip")

		if contentType := w.Header().Get("Content-Type"); contentType != "" {
			w.Header().Set("Content-Type", contentType)
		}

		gz := gzip.NewWriter(w)
		defer func() {
			if err := gz.Close(); err != nil {
				log.Printf("Error closing gzip writer: %v", err)
			}
		}()

		gzw := &gzipResponseWriter{ResponseWriter: w, Writer: gz}
		h.ServeHTTP(gzw, r)
	}
}
