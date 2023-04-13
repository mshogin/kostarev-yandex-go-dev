package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCompressHandler(t *testing.T) {
	app := &App{}
	w := httptest.NewRecorder()
	body := "http://example.com"
	r := httptest.NewRequest("POST", "/", strings.NewReader(body))

	app.CompressHandler(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("Ожидался статус код %v, получен %v", http.StatusCreated, w.Code)
	}

	responseBody := w.Body.String()
	if responseBody == "" {
		t.Errorf("Тело ответа пустое")
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
