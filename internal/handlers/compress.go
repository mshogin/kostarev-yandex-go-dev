package handlers

import (
	"fmt"
	"github.com/IKostarev/yandex-go-dev/config"
	"github.com/IKostarev/yandex-go-dev/internal/app"
	"io"
	"log"
	"net/http"
)

func CompressHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	cfg := config.LoadConfig()

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
	}

	body, _ := io.ReadAll(r.Body)

	fmt.Println(body)

	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	miniURL := app.RandomURL()
	app.SaveUrls(string(body), miniURL)

	w.WriteHeader(http.StatusCreated)
	_, err := io.WriteString(w, cfg.BaseShortURL+"/"+miniURL)
	if err != nil {
		log.Fatal(err)
	}
}
