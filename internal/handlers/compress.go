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

		m := app.SaveUrls()

		for k, v := range m {
			if url == k {
				w.Write([]byte("http://localhost:8080/" + v))
				w.WriteHeader(http.StatusCreated)
			}
		}
	}

	if r.Method == http.MethodGet {
		url := r.URL.String()

		m := app.SaveUrls()

		for k, v := range m {
			if url == ("/" + v) {
				w.Header().Add("Content-Location", k)
				w.WriteHeader(http.StatusTemporaryRedirect)
			}
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}
