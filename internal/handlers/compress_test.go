package handlers

//
//import (
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestMethodPost(t *testing.T) {
//	r := httptest.NewRequest("POST", "/", nil)
//
//	w := httptest.NewRecorder()
//
//	*App.CompressHandler(w, r)
//
//	if w.Code != http.StatusBadRequest {
//		t.Errorf("Ожидался статус код %v, получен %v", http.StatusBadRequest, w.Code)
//	}
//}
