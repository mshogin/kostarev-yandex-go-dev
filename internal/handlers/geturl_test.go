package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetURLHandler(t *testing.T) {
	app := &App{}
	w := httptest.NewRecorder()
	url := "http://ya.com"
	r := httptest.NewRequest("GET", url, nil)

	service.SaveURL(url)

	app.GetURLHandler(w, r)

	w.Header().Add("Location", url)
}

func TestGetURLHandler_BadRequest(t *testing.T) {
	app := &App{}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	app.GetURLHandler(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Ожидался статус код %v, получен %v", http.StatusBadRequest, w.Code)
	}
}
