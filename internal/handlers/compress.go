package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/app"
	"io"
	"net/http"
)

func CompressHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
	}

	body, _ := io.ReadAll(r.Body)

	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	miniURL := app.RandomURL()
	app.SaveUrls(string(body), miniURL)

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "http://localhost:8080/"+miniURL)
}
