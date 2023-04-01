package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/app"
	"net/http"
)

func CompressHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	if r.Method == http.MethodPost {
		url := r.FormValue("url")

		if len(url) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		miniURL := app.RandomURL()
		app.SaveUrls(url, miniURL)

		w.Write([]byte("http://localhost:8080/" + miniURL))
		w.WriteHeader(http.StatusCreated)
	}

	if r.Method == http.MethodGet {
		url := r.URL.String()

		m := app.GetURL(url)
		if m == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Add("Location", m)
		w.WriteHeader(http.StatusTemporaryRedirect)
	}

	w.WriteHeader(http.StatusBadRequest)
}
