package handlers

import (
	"fmt"
	"github.com/IKostarev/yandex-go-dev/internal/app"
	"net/http"
)

func GetURLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
	}

	url := r.URL.String()
	fmt.Println("r.URL.String() = ", url)
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	m := app.GetURL(url)
	fmt.Println(m)

	w.Header().Add("Location", m)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
