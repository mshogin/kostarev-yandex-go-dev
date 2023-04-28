package handlers

import (
	"bytes"
	"github.com/IKostarev/yandex-go-dev/internal/storage"
	"github.com/andybalholm/brotli"
	"github.com/go-chi/chi/v5"
	"net/http"
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

	data := []byte(m)
	//fmt.Println("data create = ", string(data))
	data, err = BrotliCompress(data)

	//fmt.Println("data posle compress = ", string(data))

	w.Header().Add("Location", string(data))

	//fmt.Println("m = ", m)

	w.Header().Set("Accept-Encoding", "gzip")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
