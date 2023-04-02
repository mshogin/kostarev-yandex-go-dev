package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/app"
	"io"
	"net/http"
)

func methodPost(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	miniURL := app.RandomURL()
	app.SaveUrls(body, miniURL)

	//w.Write([]byte("http://localhost:8080/" + miniURL))
	w.WriteHeader(http.StatusCreated)
}

func methodGet(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	m := app.GetURL(url)
	w.Header().Add("Location", string(m))
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func CompressHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	if r.Method == http.MethodPost {
		methodPost(w, r)
	} else if r.Method == http.MethodGet {
		methodGet(w, r)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
