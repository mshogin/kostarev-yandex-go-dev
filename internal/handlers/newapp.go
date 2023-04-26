package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/config"
	"io"
	"net/http"
)

type App struct {
	Config config.Config
}

type GzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *GzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
