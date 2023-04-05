package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/app"
	"net/http"
)

func GetUrlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
	}

	url := r.URL.String()
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	m := app.GetURL(url)

	w.Header().Add("Location", m)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
