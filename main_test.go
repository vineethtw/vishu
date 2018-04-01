package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	handler := healthHandler()
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Health endpoint should return 200 status code")
	}
	if w.Body.String() != "OK" {
		t.Errorf("Health endpoint has a body OK")
	}
}
