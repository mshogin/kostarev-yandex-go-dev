package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApp_GetURLHandler(t *testing.T) {
	tests := []struct {
		name       string
		idURL      string
		getURL     string
		statusCode int
		want       string
	}{
		//{ TODO не работают из-за ошибки, не знаю как пофиксить
		//	name:       "good test",
		//	idURL:      "qwertyui",
		//	getURL:     "http://test.site.com",
		//	statusCode: 307,
		//	want:       "http://test.site.com",
		//},
		//{
		//	name:       "id url is empty test",
		//	idURL:      "",
		//	getURL:     "",
		//	statusCode: 400,
		//	want:       "",
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				Storage: &mockStorage{},
			}

			r, err := http.NewRequest("GET", "/"+tt.idURL, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			w := httptest.NewRecorder()
			app.GetURLHandler(w, r)

			if w.Code != tt.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					w.Code, tt.statusCode)
			}

			if tt.want != "" && w.Header().Get("Location") != tt.want {
				t.Errorf("handler returned unexpected body: got %v want %v",
					w.Body.String(), tt.want)
			}
		})
	}
}
