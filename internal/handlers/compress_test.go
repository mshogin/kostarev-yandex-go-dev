package handlers

import (
	"github.com/IKostarev/yandex-go-dev/internal/config"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCompressHandler(t *testing.T) {
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Errorf("Не удалось прочитать config: %v", err)
	}

	app := &App{Config: cfg}
	w := httptest.NewRecorder()
	body := "http://example.com"
	r := httptest.NewRequest("POST", "/", strings.NewReader(body))

	app.CompressHandler(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("Ожидался статус код %v, получен %v", http.StatusCreated, w.Code)
	}

	responseBody := w.Body.String()
	if len(responseBody) == 0 {
		t.Errorf("Тело ответа пустое")
	}

	if err := r.Body.Close(); err != nil {
		t.Errorf("Ошибка закрытия тела запроса: %v", err)
	}
}

func TestCompressHandler_BadRequest(t *testing.T) {
	app := &App{}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)

	app.CompressHandler(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Ожидался статус код %v, получен %v", http.StatusBadRequest, w.Code)
	}
}
