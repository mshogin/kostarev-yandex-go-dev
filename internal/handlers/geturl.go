package handlers

import (
	"bytes"
	"fmt"
	"github.com/IKostarev/yandex-go-dev/internal/storage"
	"github.com/andybalholm/brotli"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

func BrotliCompress(data []byte) ([]byte, error) {
	var buf bytes.Buffer

	w := brotli.NewWriterLevel(&buf, brotli.BestCompression)
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

func (a *App) GetURLHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/plain")

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

	accept := strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")
	content := strings.Contains(r.Header.Get("Content-Encoding"), "gzip")

	if accept || content {
		fmt.Println(accept)
		fmt.Println(content)
		data, _ := BrotliCompress([]byte(m)) //TODO обработать ошибку
		fmt.Println("m = ", m)
		fmt.Println("data = ", data)
		fmt.Println("data = ", string(data))
		w.Header().Add("Location", string(data))
		w.Header().Set("Accept-Encoding", "gzip")
	} else {
		w.Header().Add("Location", m)
		fmt.Println("m = ", m)
	}
	w.WriteHeader(http.StatusTemporaryRedirect)
}
