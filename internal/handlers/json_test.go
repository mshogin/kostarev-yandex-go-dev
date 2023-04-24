package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/IKostarev/yandex-go-dev/internal/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSONHandler(t *testing.T) {
	app := &App{}

	url := "https://example.com"
	reqBody := []byte(`{"url": "` + url + `"}`)
	req, err := http.NewRequest("POST", "/api/shorten", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	app.JSONHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}

	var responseMap map[string]string

	err = json.Unmarshal(w.Body.Bytes(), &responseMap)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	url = storage.SaveURL(url)
	responseMap["result"] = url

	expectedURL := app.Config.BaseShortURL + url
	if responseMap["result"] != expectedURL {
		t.Errorf("Expected response body %s, but got %s", expectedURL, responseMap["result"])
	}
}
